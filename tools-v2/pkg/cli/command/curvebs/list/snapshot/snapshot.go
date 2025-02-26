package snapshot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	cmderror "github.com/opencurve/curve/tools-v2/internal/error"
	cobrautil "github.com/opencurve/curve/tools-v2/internal/utils"
	basecmd "github.com/opencurve/curve/tools-v2/pkg/cli/command"
	"github.com/opencurve/curve/tools-v2/pkg/config"
	"github.com/opencurve/curve/tools-v2/pkg/output"
	"github.com/spf13/cobra"
)

const (
	snapshotExample = `$ curve bs list snapshot`
)

type Response struct {
	Code       string          `json:"Code"`
	Message    string          `json:"Message"`
	RequestId  string          `json:"RequestId"`
	TotalCount int             `json:"TotalCount"`
	SnapShots  []*SnapShotInfo `json:"Snapshots"`
}

type SnapShotInfo struct {
	File       string `json:"File"`
	FileLength int    `json:"FileLength"`
	Name       string `json:"Name"`
	Progress   int    `json:"Progress"`
	SeqNum     int    `json:"SeqNum"`
	Status     int    `json:"Status"`
	Time       int    `json:"Time"`
	UUID       string `json:"UUID"`
	User       string `json:"User"`
}

type SnapShotCommand struct {
	basecmd.FinalCurveCmd
	snapshotAddrs []string
	timeout       time.Duration

	user string
	file string
	uuid string
}

var _ basecmd.FinalCurveCmdFunc = (*SnapShotCommand)(nil)

func NewSnapShotCommand() *cobra.Command {
	return NewListSnapShotCommand().Cmd
}

func NewListSnapShotCommand() *SnapShotCommand {
	snapShotCommand := &SnapShotCommand{
		FinalCurveCmd: basecmd.FinalCurveCmd{
			Use:     "snapshot",
			Short:   "list snapshot information in curvebs",
			Example: snapshotExample,
		},
	}

	basecmd.NewFinalCurveCli(&snapShotCommand.FinalCurveCmd, snapShotCommand)
	return snapShotCommand
}

func (sCmd *SnapShotCommand) AddFlags() {
	config.AddBsSnapshotCloneFlagOption(sCmd.Cmd)
	config.AddHttpTimeoutFlag(sCmd.Cmd)
	config.AddBsUserOptionFlag(sCmd.Cmd)
	config.AddBsSnapshotIDOptionFlag(sCmd.Cmd)
	config.AddBsPathOptionFlag(sCmd.Cmd)
}

func (sCmd *SnapShotCommand) Init(cmd *cobra.Command, args []string) error {
	snapshotAddrs, err := config.GetBsSnapshotAddrSlice(sCmd.Cmd)
	if err.TypeCode() != cmderror.CODE_SUCCESS || len(snapshotAddrs) == 0 {
		return err.ToError()
	}
	sCmd.snapshotAddrs = snapshotAddrs
	sCmd.timeout = config.GetFlagDuration(sCmd.Cmd, config.HTTPTIMEOUT)
	sCmd.user = config.GetBsFlagString(sCmd.Cmd, config.CURVEBS_USER)
	sCmd.file = config.GetBsFlagString(sCmd.Cmd, config.CURVEBS_PATH)
	sCmd.uuid = config.GetBsFlagString(sCmd.Cmd, config.CURVEBS_SNAPSHOT_ID)
	header := []string{
		cobrautil.ROW_SNAPSHOT_ID,
		cobrautil.ROW_SNAPSHOT_NAME,
		cobrautil.ROW_USER,
		cobrautil.ROW_STATUS,
		cobrautil.ROW_SNAPSHOT_SEQNUM,
		cobrautil.ROW_FILE_LENGTH,
		cobrautil.ROW_PROGRESS,
		cobrautil.ROW_CREATE_TIME,
		cobrautil.ROW_FILE,
	}
	sCmd.SetHeader(header)
	sCmd.TableNew.SetAutoMergeCellsByColumnIndex(cobrautil.GetIndexSlice(
		sCmd.Header, []string{cobrautil.ROW_FILE},
	))
	return nil
}

func (sCmd *SnapShotCommand) Print(cmd *cobra.Command, args []string) error {
	return output.FinalCmdOutput(&sCmd.FinalCurveCmd, sCmd)
}

func (sCmd *SnapShotCommand) RunCommand(cmd *cobra.Command, args []string) error {
	params := map[string]any{
		cobrautil.QueryAction:  cobrautil.ActionGetFileSnapshotList,
		cobrautil.QueryVersion: cobrautil.Version,
		cobrautil.QueryUser:    sCmd.user,
		cobrautil.QueryFile:    sCmd.file,
		cobrautil.QueryLimit:   cobrautil.Limit,
		cobrautil.QueryOffset:  cobrautil.Offset,
	}
	if sCmd.uuid != "*" {
		params[cobrautil.QueryUUID] = sCmd.uuid
	}
	snapshotsInfo, err := ListSnapShot(sCmd.snapshotAddrs, sCmd.timeout, params)
	if err != nil {
		sCmd.Error = err
		return sCmd.Error.ToError()
	}
	rows := make([]map[string]string, 0)
	for _, item := range snapshotsInfo {
		row := make(map[string]string)
		row[cobrautil.ROW_SNAPSHOT_ID] = item.UUID
		row[cobrautil.ROW_SNAPSHOT_NAME] = item.Name
		row[cobrautil.ROW_USER] = item.User
		row[cobrautil.ROW_FILE] = item.File
		row[cobrautil.ROW_STATUS] = fmt.Sprintf("%d", item.Status)
		row[cobrautil.ROW_SNAPSHOT_SEQNUM] = fmt.Sprintf("%d", item.SeqNum)
		row[cobrautil.ROW_FILE_LENGTH] = fmt.Sprintf("%d", item.FileLength)
		row[cobrautil.ROW_PROGRESS] = fmt.Sprintf("%d", item.Progress)
		row[cobrautil.ROW_CREATE_TIME] = time.Unix(int64(item.Time/1000000), 0).Format("2006-01-02 15:04:05")
		rows = append(rows, row)
	}
	list := cobrautil.ListMap2ListSortByKeys(rows, sCmd.Header, []string{cobrautil.ROW_FILE, cobrautil.ROW_SNAPSHOT_NAME, cobrautil.ROW_SNAPSHOT_ID})
	sCmd.TableNew.AppendBulk(list)
	sCmd.Result = rows
	sCmd.Error = cmderror.Success()
	return nil
}

func (sCmd *SnapShotCommand) ResultPlainOutput() error {
	return output.FinalCmdOutputPlain(&sCmd.FinalCurveCmd)
}

func ListSnapShot(addrs []string, timeout time.Duration, params map[string]any) ([]*SnapShotInfo, *cmderror.CmdError) {
	var snapshotsInfo []*SnapShotInfo
	for {
		var resp Response
		subUri := cobrautil.NewSubUri(params)
		metric := basecmd.NewMetric(addrs, subUri, timeout)

		result, err := basecmd.QueryMetric(metric)
		if err.TypeCode() != cmderror.CODE_SUCCESS {
			return snapshotsInfo, err
		}
		if err := json.Unmarshal([]byte(result), &resp); err != nil {
			retErr := cmderror.ErrUnmarshalJson()
			retErr.Format(err.Error())
			return snapshotsInfo, retErr
		}
		if len(resp.SnapShots) == 0 || resp.SnapShots == nil {
			return snapshotsInfo, nil
		}
		snapshotsInfo = append(snapshotsInfo, resp.SnapShots...)
		offsetValue, _ := strconv.Atoi(params[cobrautil.QueryOffset].(string))
		limitValue, _ := strconv.Atoi(params[cobrautil.QueryLimit].(string))
		params[cobrautil.QueryOffset] = strconv.Itoa(offsetValue + limitValue)
	}
}
