# curve tool usage

A tool for CurveFS & CurveBs.

- [curve tool usage](#curve-tool-usage)
  - [How to use curve tool](#how-to-use-curve-tool)
    - [Install](#install)
    - [Introduction](#introduction)
  - [Command](#command)
    - [version](#version)
    - [completion](#completion)
    - [upgrade](#upgrade)
    - [fs](#fs)
      - [check](#check)
        - [check copyset](#check-copyset)
      - [create](#create)
        - [create fs](#create-fs)
        - [create topology](#create-topology)
      - [delete](#delete)
        - [delete fs](#delete-fs)
      - [list](#list)
        - [list copyset](#list-copyset)
        - [list fs](#list-fs)
        - [list mountpoint](#list-mountpoint)
        - [list partition](#list-partition)
        - [list topology](#list-topology)
      - [query](#query)
        - [query copyset](#query-copyset)
        - [query fs](#query-fs)
        - [query inode](#query-inode)
        - [query metaserver](#query-metaserver)
        - [query partition](#query-partition)
      - [status](#status)
        - [status mds](#status-mds)
        - [status metaserver](#status-metaserver)
        - [status etcd](#status-etcd)
        - [status copyset](#status-copyset)
        - [status cluster](#status-cluster)
      - [umount](#umount)
        - [umount fs](#umount-fs)
      - [usage](#usage)
        - [usage inode](#usage-inode)
        - [usage metadata](#usage-metadata)
      - [warmup](#warmup)
        - [warmup add](#warmup-add)
        - [warmup cancel](#warmup-cancel)
        - [warmup list](#warmup-list)
        - [warmup query](#warmup-query)
    - [bs](#bs)
      - [list](#list-1)
          - [list logical-pool](#list-logical-pool)
          - [list server](#list-server)
          - [list client](#list-client)
          - [list dir](#list-dir)
          - [list space](#list-space)
          - [list chunkserver](#list-chunkserver)
          - [list scan-status](#list-scan-status)
          - [list may-broken-vol](#list-may-broken-vol)
          - [list snapshot](#list-snapshot)
      - [clean-recycle](#clean-recycle)
      - [query](#query-1)
          - [query file](#query-file)
          - [query chunk](#query-chunk)
          - [query segment](#query-segment)
        - [query scan-status](#query-scan-status)
      - [status](#status-1)
        - [status etcd](#status-etcd-1)
        - [status mds](#status-mds-1)
        - [status client](#status-client)
        - [status snapshotserver](#status-snapshotserver)
        - [status chunkserver](#status-chunkserver)
        - [status copyset](#status-copyset-1)
      - [delete](#delete-1)
        - [delete peer](#delete-peer)
        - [delete volume](#delete-volume)
          - [delete volume clone](#delete-volume-clone)
          - [delete volume recover](#delete-volume-recover)
      - [update](#update)
        - [update peer](#update-peer)
        - [update leader](#update-leader)
        - [update file](#update-file)
        - [update throttle](#update-throttle)
        - [update scan-state](#update-scan-state)
        - [update copyset availflag](#update-copyset-availflag)
        - [update leader-schedule](#update-leader-schedule)
      - [create](#create-1)
        - [create file](#create-file)
        - [create dir](#create-dir)
      - [check](#check-1)
        - [check copyset](#check-copyset-1)
        - [check chunkserver](#check-chunkserver)
        - [check server](#check-server)
        - [check consistency](#check-consistency)
      - [snapshot](#snapshot)
        - [snapshot copyset](#snapshot-copyset)
  - [Comparison of old and new commands](#comparison-of-old-and-new-commands)
    - [curve fs](#curve-fs)
    - [curve bs](#curve-bs)

## How to use curve tool

### Install

install curve tool

```bash
wget https://curve-tool.nos-eastchina1.126.net/release/curve-latest
chmod +x curve-latest
mv curve-latest /usr/bin/curve
```

set configure file

```bash
wget https://raw.githubusercontent.com/opencurve/curve/master/tools-v2/pkg/config/curve.yaml
```

or

```
wget https://curve-tool.nos-eastchina1.126.net/config/curve.yaml
```

Please modify the `mdsAddr, mdsDummyAddr, etcdAddr` under `curvefs/bs` in the template.yaml file as required

```bash
mv curve.yaml ~/.curve/curve.yaml
```

or

```bash
mv curve.yaml /etc/curve/curve.yaml
```

### Introduction

Here's how to use the tool

```bash
curve COMMAND [options]
```

When you are not sure how to use a command, --help can give you an example of use:

```bash
curve COMMAND --help
```

For example:

```bash
curve fs status mds --help
Usage:  curve fs status mds [flags]

get the inode usage of curvefs

Flags:
  -c, --conf string            config file (default is $HOME/.curve/curve.yaml or /etc/curve/curve.yaml)
  -f, --format string          Output format (json|plain) (default "plain")
  -h, --help                   Print usage
      --httptimeout duration   http timeout (default 500ms)
      --mdsaddr string         mds address, should be like 127.0.0.1:6700,127.0.0.1:6701,127.0.0.1:6702
      --mdsdummyaddr string    mds dummy address, should be like 127.0.0.1:7700,127.0.0.1:7701,127.0.0.1:7702
      --showerror              display all errors in command

Examples:
$ curve fs status mds
```

In addition, this tool reads the configuration from `$HOME/.curve/curve.yaml` or `/etc/curve/curve.yaml` by default,
and can be specified by `--conf` or `-c`.

## Command

### version

show the version of curve tool

Usage:

```shell
curve --version
```

Output:

```shell
curve v1.2
```

### completion

generate curve bash/zsh/fish completion script

Take bash as an example:
```bash
curve completion bash > /etc/bash_completion.d/curve
```

If you perform completion you get the following error:
```bash
curve [tab]bash: _get_comp_words_by_ref: command not found
```

You need to install bash-completion:
```bash
sudo apt install bash-completion
source /usr/share/bash-completion/bash_completion
```

### upgrade

Upgrade curve to latest version

Usage:

```shell
curve upgrade
```

If you want to upgrade to the specified version(ee500438):

```bash
CURVE_VERSION=ee500438 curve upgrade
```

If you are in an offline environment, you can also set up a temporary http server to provide upgrade services for other machines:

```bash
# 192.168.1.1
echo "123456" > __version
cp curve curve-123456
python3 -m http.server 1234
```

In this way, it can be upgraded on other machines

```bash
CURVE_BASE_URL=http://192.168.1.1:1234/ curve upgrade
```

### fs

#### check

##### check copyset

check copysets health in curvefs

Usage:

```shell
curve fs check copyset --copysetid 1 --poolid 1
```

Output:

```shell
+------------+-----------+--------+--------+---------+
| COPYSETKEY | COPYSETID | POOLID | STATUS | EXPLAIN |
+------------+-----------+--------+--------+---------+
| 4294967297 |         1 |      1 | ok     |         |
+------------+-----------+--------+--------+---------+
```

#### create

##### create fs

create a fs in curvefs

Usage:

```shell
curve fs create fs --fsname test3
```

Output:

```shell
+--------+-------------------------------------------+
| FSNAME |                  RESULT                   |
+--------+-------------------------------------------+
| test3  | fs exist, but s3 info is not inconsistent |
+--------+-------------------------------------------+
```

##### create topology

create curvefs topology

Usage:

```shell
curve fs create topology --clustermap topology.json
```

Output:

```shell
+-------------------+--------+-----------+--------+
|       NAME        |  TYPE  | OPERATION | PARENT |
+-------------------+--------+-----------+--------+
| pool2             | pool   | del       |        |
+-------------------+--------+           +--------+
| zone4             | zone   |           | pool2  |
+-------------------+--------+           +--------+
| **.***.***.**_3_0 | server |           | zone4  |
+-------------------+--------+-----------+--------+
```

#### delete

##### delete fs

delete a fs from curvefs

Usage:

```shell
curve fs delete fs --fsname test1
WARNING:Are you sure to delete fs test1?
please input [test1] to confirm: test1
```

Output:

```shell
+--------+-------------------------------------+
| FSNAME |               RESULT                |
+--------+-------------------------------------+
| test1  | delete fs failed!, error is FS_BUSY |
+--------+-------------------------------------+
```

#### list

##### list copyset

list all copyset info of the curvefs

Usage:

```shell
curve fs list copyset
```

Output:

```shell
+------------+-----------+--------+-------+--------------------------------+------------+
|    KEY     | COPYSETID | POOLID | EPOCH |           LEADERPEER           | PEERNUMBER |
+------------+-----------+--------+-------+--------------------------------+------------+
| 4294967302 | 6         | 1      | 2     | id:1                           | 3          |
|            |           |        |       | address:"**.***.***.**:6801:0" |            |
+------------+-----------+        +-------+                                +------------+
| 4294967303 | 7         |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+                                +------------+
| 4294967304 | 8         |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+                                +------------+
| 4294967307 | 11        |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+--------------------------------+------------+
| 4294967297 | 1         |        | 1     | id:2                           | 3          |
|            |           |        |       | address:"**.***.***.**:6802:0" |            |
+------------+-----------+        +-------+                                +------------+
| 4294967301 | 5         |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+                                +------------+
| 4294967308 | 12        |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+--------------------------------+------------+
| 4294967298 | 2         |        | 1     | id:3                           | 3          |
|            |           |        |       | address:"**.***.***.**:6800:0" |            |
+------------+-----------+        +-------+                                +------------+
| 4294967299 | 3         |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+                                +------------+
| 4294967300 | 4         |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+                                +------------+
| 4294967305 | 9         |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+        +-------+                                +------------+
| 4294967306 | 10        |        | 1     |                                | 3          |
|            |           |        |       |                                |            |
+------------+-----------+--------+-------+--------------------------------+------------+
```

##### list fs

list all fs info in the curvefs

Usage:

```shell
curve fs list fs
```

Output:

```shell
+----+-------+--------+--------------+-----------+---------+----------+-----------+----------+
| ID | NAME  | STATUS |   CAPACITY   | BLOCKSIZE | FSTYPE  | SUMINDIR |   OWNER   | MOUNTNUM |
+----+-------+--------+--------------+-----------+---------+----------+-----------+----------+
| 2  | test1 | INITED | 107374182400 | 1048576   | TYPE_S3 | false    | anonymous | 1        |
+----+-------+--------+--------------+-----------+         +----------+           +----------+
| 3  | test3 | INITED | 107374182400 | 1048576   |         | false    |           | 0        |
+----+-------+--------+--------------+-----------+---------+----------+-----------+----------+
```

##### list mountpoint

list all mountpoint of the curvefs

Usage:

```shell
curve fs list mountpoint
```

Output:

```shell
+------+--------+-------------------------------------------------------------------+
| FSID | FSNAME |                            MOUNTPOINT                             |
+------+--------+-------------------------------------------------------------------+
| 2    | test1  | siku-QiTianM420-N000:9002:/curvefs/client/mnt/home/siku/temp/mnt1 |
+      +        +-------------------------------------------------------------------+
|      |        | siku-QiTianM420-N000:9003:/curvefs/client/mnt/home/siku/temp/mnt2 |
+------+--------+-------------------------------------------------------------------+
```

##### list partition

list partition in curvefs by fsid

Usage:

```shell
curve fs list partition
```

Output:

```shell
+-------------+------+--------+-----------+----------+----------+-----------+
| PARTITIONID | FSID | POOLID | COPYSETID |  START   |   END    |  STATUS   |
+-------------+------+--------+-----------+----------+----------+-----------+
| 14          | 2    | 1      | 10        | 1048676  | 2097351  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 20          |      |        |           | 7340732  | 8389407  | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 13          |      |        | 11        | 0        | 1048675  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 16          |      |        |           | 3146028  | 4194703  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 22          |      |        |           | 9438084  | 10486759 | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 21          |      |        | 5         | 8389408  | 9438083  | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 23          |      |        | 7         | 10486760 | 11535435 | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 24          |      |        |           | 11535436 | 12584111 | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 15          |      |        | 8         | 2097352  | 3146027  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 18          |      |        |           | 5243380  | 6292055  | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 17          |      |        | 9         | 4194704  | 5243379  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 19          |      |        |           | 6292056  | 7340731  | READWRITE |
+-------------+------+        +-----------+----------+----------+-----------+
| 26          | 3    |        | 2         | 1048676  | 2097351  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 30          |      |        |           | 5243380  | 6292055  | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 34          |      |        | 3         | 9438084  | 10486759 | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 29          |      |        | 4         | 4194704  | 5243379  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 32          |      |        |           | 7340732  | 8389407  | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 35          |      |        | 5         | 10486760 | 11535435 | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 27          |      |        |           | 2097352  | 3146027  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 33          |      |        |           | 8389408  | 9438083  | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 25          |      |        | 6         | 0        | 1048675  | READWRITE |
+-------------+      +        +           +----------+----------+-----------+
| 36          |      |        |           | 11535436 | 12584111 | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 28          |      |        | 8         | 3146028  | 4194703  | READWRITE |
+-------------+      +        +-----------+----------+----------+-----------+
| 31          |      |        | 9         | 6292056  | 7340731  | READWRITE |
+-------------+------+--------+-----------+----------+----------+-----------+
```

##### list topology

list the topology of the curvefs

Usage:

```shell
curve fs list topology
```

Output:

```shell
+----+------------+--------------------+------------+-----------------------+
| ID |    TYPE    |        NAME        | CHILDTYPE  |       CHILDLIST       |
+----+------------+--------------------+------------+-----------------------+
| 1  | pool       | pool1              | zone       | zone3 zone2 zone1     |
+----+------------+--------------------+------------+-----------------------+
| 3  | zone       | zone3              | server     | **.***.***.**_2_0     |
+----+            +--------------------+            +-----------------------+
| 2  |            | zone2              |            | **.***.***.**_1_0     |
+----+            +--------------------+            +-----------------------+
| 1  |            | zone1              |            | **.***.***.**_0_0     |
+----+------------+--------------------+------------+-----------------------+
| 3  | server     | **.***.***.**_2_0  | metaserver | curvefs-metaserver.2  |
+----+            +--------------------+            +-----------------------+
| 2  |            | **.***.***.**_1_0  |            | curvefs-metaserver.1  |
+----+            +--------------------+            +-----------------------+
| 1  |            | **.***.***.**_0_0  |            | curvefs-metaserver.3  |
+----+------------+--------------------+------------+-----------------------+
| 3  | metaserver | curvefs-metaserver |            |                       |
+----+            +--------------------+------------+-----------------------+
| 2  |            | curvefs-metaserver |            |                       |
+----+            +--------------------+------------+-----------------------+
| 1  |            | curvefs-metaserver |            |                       |
+----+------------+--------------------+------------+-----------------------+
```

#### query

##### query copyset

query copysets in curvefs

Usage:

```shell
curve fs query copyset --copysetid 1 --poolid 1
```

Output:

```shell
+------------+-----------+--------+--------------------------------------+-------+
| copysetKey | copysetId | poolId |              leaderPeer              | epoch |
+------------+-----------+--------+--------------------------------------+-------+
| 4294967297 |     1     |   1    | id:2  address:"**.***.***.**:6802:0" |   1   |
+------------+-----------+--------+--------------------------------------+-------+
```

##### query fs

query fs in curvefs by fsname or fsid

Usage:

```shell
curve fs query fs --fsname test1
```

Output:

```shell
+----+-------+--------+--------------+-----------+---------+----------+-----------+----------+
| id | name  | status |   capacity   | blocksize | fsType  | sumInDir |   owner   | mountNum |
+----+-------+--------+--------------+-----------+---------+----------+-----------+----------+
| 2  | test1 | INITED | 107374182400 |  1048576  | TYPE_S3 |  false   | anonymous |    2     |
+----+-------+--------+--------------+-----------+---------+----------+-----------+----------+
```

##### query inode

query the inode of fs

Usage:

```shell
curve fs query inode --fsid 2 --inodeid 5243380
```

Output:

```shell
+-------+----------+-----------+---------+-------+--------+
| fs id | inode id |  length   |  type   | nlink | parent |
+-------+----------+-----------+---------+-------+--------+
|   2   | 5243380  | 352321536 | TYPE_S3 |   1   |  [1]   |
+-------+----------+-----------+---------+-------+--------+
```

##### query metaserver

query metaserver in curvefs by metaserverid or metaserveraddr

Usage:

```shell
curve fs query metaserver --metaserveraddr **.***.***.**:6801,**.***.***.**:6802
```

Output:

```shell
+----+--------------------+--------------------+--------------------+-------------+
| id |      hostname      |    internalAddr    |    externalAddr    | onlineState |
+----+--------------------+--------------------+--------------------+-------------+
| 1  | curvefs-metaserver | **.***.***.**:6801 | **.***.***.**:6801 |   ONLINE    |
| 2  | curvefs-metaserver | **.***.***.**:6802 | **.***.***.**:6802 |   ONLINE    |
+----+--------------------+--------------------+--------------------+-------------+
```

##### query partition

query the copyset of partition

Usage:

```shell
curve fs query partition --partitionid 14
```

Output:

```shell
+----+--------+-----------+--------+----------------------+
| id | poolId | copysetId | peerId |       peerAddr       |
+----+--------+-----------+--------+----------------------+
| 14 |   1    |    10     |   1    | **.***.***.**:6801:0 |
| 14 |   1    |    10     |   2    | **.***.***.**:6802:0 |
| 14 |   1    |    10     |   3    | **.***.***.**:6800:0 |
+----+--------+-----------+--------+----------------------+
```

#### status

##### status mds

get status of mds

Usage:

```shell
curve fs status mds
```

Output:

```shell
+--------------------+--------------------+----------------+----------+
|        addr        |     dummyAddr      |    version     |  status  |
+--------------------+--------------------+----------------+----------+
| **.***.***.**:6700 | **.***.***.**:7700 | 8fc48476+debug | follower |
| **.***.***.**:6701 | **.***.***.**:7701 | 8fc48476+debug | follower |
| **.***.***.**:6702 | **.***.***.**:7702 | 8fc48476+debug |  leader  |
+--------------------+--------------------+----------------+----------+
```

##### status metaserver

get status of metaserver

Usage:

```shell
curve fs status metaserver
```

Output:

```shell
+--------------------+--------------------+----------------+--------+
|    externalAddr    |    internalAddr    |    version     | status |
+--------------------+--------------------+----------------+--------+
| **.***.***.**:6800 | **.***.***.**:6800 | 8fc48476+debug | online |
| **.***.***.**:6802 | **.***.***.**:6802 | 8fc48476+debug | online |
| **.***.***.**:6801 | **.***.***.**:6801 | 8fc48476+debug | online |
+--------------------+--------------------+----------------+--------+
```

##### status etcd

get status of etcd

Usage:

```shell
curve fs status etcd
```

Output:

```shell
+---------------------+---------+----------+
|        addr         | version |  status  |
+---------------------+---------+----------+
| **.***.***.**:23790 | 3.4.10  | follower |
| **.***.***.**:23791 | 3.4.10  | follower |
| **.***.***.**:23792 | 3.4.10  |  leader  |
+---------------------+---------+----------+
```

##### status copyset

get status of copyset

Usage:

```shell
curve fs status copyset
```

Output:

```shell
+------------+-----------+--------+--------+---------+
| copysetKey | copysetId | poolId | status | explain |
+------------+-----------+--------+--------+---------+
| 4294967297 |     1     |   1    |   ok   |         |
| 4294967298 |     2     |   1    |   ok   |         |
| 4294967299 |     3     |   1    |   ok   |         |
| 4294967300 |     4     |   1    |   ok   |         |
| 4294967301 |     5     |   1    |   ok   |         |
| 4294967302 |     6     |   1    |   ok   |         |
| 4294967303 |     7     |   1    |   ok   |         |
| 4294967304 |     8     |   1    |   ok   |         |
| 4294967305 |     9     |   1    |   ok   |         |
| 4294967306 |    10     |   1    |   ok   |         |
| 4294967307 |    11     |   1    |   ok   |         |
| 4294967308 |    12     |   1    |   ok   |         |
+------------+-----------+--------+--------+---------+
```

##### status cluster

get status of cluster

Usage:

```shell
curve fs status cluster
```

Output:

```shell
etcd:
+---------------------+---------+----------+
|        addr         | version |  status  |
+---------------------+---------+----------+
| **.***.***.**:23790 | 3.4.10  | follower |
| **.***.***.**:23791 | 3.4.10  | follower |
| **.***.***.**:23792 | 3.4.10  |  leader  |
+---------------------+---------+----------+

mds:
+--------------------+--------------------+----------------+----------+
|        addr        |     dummyAddr      |    version     |  status  |
+--------------------+--------------------+----------------+----------+
| **.***.***.**:6700 | **.***.***.**:7700 | 8fc48476+debug | follower |
| **.***.***.**:6701 | **.***.***.**:7701 | 8fc48476+debug | follower |
| **.***.***.**:6702 | **.***.***.**:7702 | 8fc48476+debug |  leader  |
+--------------------+--------------------+----------------+----------+

meataserver:
+--------------------+--------------------+----------------+--------+
|    externalAddr    |    internalAddr    |    version     | status |
+--------------------+--------------------+----------------+--------+
| **.***.***.**:6800 | **.***.***.**:6800 | 8fc48476+debug | online |
| **.***.***.**:6802 | **.***.***.**:6802 | 8fc48476+debug | online |
| **.***.***.**:6801 | **.***.***.**:6801 | 8fc48476+debug | online |
+--------------------+--------------------+----------------+--------+

copyset:
+------------+-----------+--------+--------+---------+
| copysetKey | copysetId | poolId | status | explain |
+------------+-----------+--------+--------+---------+
| 4294967297 |     1     |   1    |   ok   |         |
| 4294967298 |     2     |   1    |   ok   |         |
| 4294967299 |     3     |   1    |   ok   |         |
| 4294967300 |     4     |   1    |   ok   |         |
| 4294967301 |     5     |   1    |   ok   |         |
| 4294967302 |     6     |   1    |   ok   |         |
| 4294967303 |     7     |   1    |   ok   |         |
| 4294967304 |     8     |   1    |   ok   |         |
| 4294967305 |     9     |   1    |   ok   |         |
| 4294967306 |    10     |   1    |   ok   |         |
| 4294967307 |    11     |   1    |   ok   |         |
| 4294967308 |    12     |   1    |   ok   |         |
+------------+-----------+--------+--------+---------+

Cluster health is:  ok
```

#### umount

##### umount fs

umount fs from the curvefs cluster

Usage:

```shell
curve fs umount fs --fsname test1 --mountpoint siku-QiTianM420-N000:9002:/curvefs/client/mnt/home/siku/temp/mnt1
```

Output:

```shell
+--------+-------------------------------------------------------------------+---------+
| fsName |                            mountpoint                             | result  |
+--------+-------------------------------------------------------------------+---------+
| test1  | siku-QiTianM420-N000:9003:/curvefs/client/mnt/home/siku/temp/mnt2 | success |
+--------+-------------------------------------------------------------------+---------+
```

#### usage

##### usage inode

get the inode usage of curvefs

Usage:

```shell
curve fs usage inode
```

Output:

```shell
+------+----------------+-----+
| fsId |     fsType     | num |
+------+----------------+-----+
|  2   |   inode_num    |  3  |
|  2   | type_directory |  1  |
|  2   |   type_file    |  0  |
|  2   |    type_s3     |  2  |
|  2   | type_sym_link  |  0  |
|  3   |   inode_num    |  1  |
|  3   | type_directory |  1  |
|  3   |   type_file    |  0  |
|  3   |    type_s3     |  0  |
|  3   | type_sym_link  |  0  |
+------+----------------+-----+
```

##### usage metadata

get the usage of metadata in curvefs

Usage:

```shell
curve fs usage metadata
```

Output:

```shell
+--------------------+---------+---------+---------+
|   metaserverAddr   |  total  |  used   |  left   |
+--------------------+---------+---------+---------+
| **.***.***.**:6800 | 2.0 TiB | 182 GiB | 1.8 TiB |
| **.***.***.**:6802 | 2.0 TiB | 182 GiB | 1.8 TiB |
| **.***.***.**:6801 | 2.0 TiB | 182 GiB | 1.8 TiB |
+--------------------+---------+---------+---------+
```

#### warmup

When using object storage with CurveFS, our data is stored remotely on platforms like MinIO or AWS S3.
Each time we access a file within the CurveFS mount, we have to fetch the content from the remote platform.
This process involves overhead like network connections, file writing, and memory access.

To alleviate the wait when users require a file, the warmup tool allows users to pre-fetch files of interest in advance.
This way, the needed files are readily available when users need them.

#### warmup add

Initiate warming up for either a file or directory, or provide a single file (file list) containing a list of files and directories you want to warm up.

Usage:

```shell
curve fs warmup add [--filelist] <target>
curve fs warmup add /mnt/curvefs/warmup
curve fs warmup add --filelist /mnt/curvefs/warmup.list
```

> `curve fs warmup add /mnt/curvefs/warmup` will warmup a file(directory).

> `curve fs warmup add --filelist /mnt/curvefs/warmup.list` will warmup a filelist.

#### warmup cancel

Cancel a warmup job that's currently in progress for either a single file or directory, or for a list of files and directories contained within a file list.

Usage:

```shell
curve fs warmup cancel [--filelist] <target>
curve fs warmup cancel /mnt/curvefs/warmup
curve fs warmup cancel --filelist /mnt/curvefs/warmup.list
```

> `curve fs warmup cancel /mnt/curvefs/warmup` will cancel a warmup job running on the speicfied file(directory).

> `curve fs warmup cancel --filelist /mnt/curvefs/warmup.list` will cancel a warmup job running on the speicfied filelist.

#### warmup list

Display a list of all currently active warmup jobs initiated within the specified CurveFS filesystem.

Usage:

```shell
curve fs warmup list <curvefs mount path>
curve fs warmup list /mnt/curvefs
```

> `curve fs warmup list /mnt/curvefs` will list out all the running warmup jobs within the specified CurveFS filesystem mount path.

#### warmup query

Check the status of a warmup job for either a single file or directory, or for a list of files and directories contained within a file list.

Usage:

```shell
curve fs warmup query <target>
curve fs warmup query /mnt/curvefs/warmup
```

> `curve fs warmup query /mnt/curvefs/warmup` will display the warmup job progress running on the specified target.

### bs

#### list

###### list logical-pool

list all logical pool information

Usage:

```bash
curve bs list logical-pool
```

Output:

```bash
+----+-------+-----------+----------+-------+------+--------+------+--------+---------+
| ID | NAME  | PHYPOOLID |   TYPE   | ALLOC | SCAN | TOTAL  | USED |  LEFT  | RECYCLE |
+----+-------+-----------+----------+-------+------+--------+------+--------+---------+
| 1  | pool1 | 1         | PAGEFILE | ALLOW | true | 44 GiB | 0 B  | 44 GiB | 0 B     |
+----+-------+-----------+----------+-------+------+--------+------+--------+---------+
```

###### list server

list all server information in curvebs

Usage:

```bash
curve bs list server
```

Output:

```bash
+----+---------------------+------+---------+-------------------+-------------------+
| ID |      HOSTNAME       | ZONE | PHYPOOL |   INTERNALADDR    |   EXTERNALADDR    |
+----+---------------------+------+---------+-------------------+-------------------+
| 1  | ***************_0_0 | 1    | 1       | **.***.**.**:**** | **.***.**.**:**** |
+----+---------------------+------+         +-------------------+-------------------+
| 2  | ***************_1_0 | 2    |         | **.***.**.**:**** | **.***.**.**:**** |
+----+---------------------+------+         +-------------------+-------------------+
| 3  | ***************_2_0 | 3    |         | **.***.**.**:**** | **.***.**.**:**** |
+----+---------------------+------+---------+-------------------+-------------------+
```

###### list client

list all client information in curvebs

```bash
curve bs list client
```

Output:

```bash
+------------+------+
|     IP     | PORT |
+------------+------+
| 172.17.0.2 | 9000 |
+------------+------+
```

###### list dir

list dir information in curvebs

```bash
curve bs list dir --path /
```

Output:

```bash
+------+-------------+----------+-----------------+------------+---------------------+---------------+-------------+
|  ID  |  FILENAME   | PARENTID |    FILETYPE     |   OWNER    |        CTIME        | ALLOCATEDSIZE |  FILESIZE   |
+------+-------------+----------+-----------------+------------+---------------------+---------------+-------------+
| 1    | /RecycleBin | 0        | INODE_DIRECTORY | root       | 2022-11-12 16:38:25 |      0 B      |     0 B     |
+------+-------------+----------+-----------------+------------+---------------------+---------------+-------------+
```

###### list space

show curvebs all disk type space, include total space and used space

```bash
curve bs list space
```

Output:

```bash
+----------+---------+---------+---------+------------+---------+
|   TYPE   | TOTAL   |  USED   |  LEFT   | RECYCLABLE | CREATED |
+----------+---------+---------+---------+------------+---------+
| physical | *** GiB | *** GiB | *** GiB | -          | -       |
+----------+---------+---------+---------+------------+---------+
| logical  | *** GiB | *** GiB | *** GiB | *** GiB    | *** GiB |
+----------+---------+---------+---------+------------+---------+
```

###### list chunkserver

list chunkserver information in curvebs

```bash
curve bs list chunkserver
```

Output:

```bash        
+----+------+-----------+------+-----------+------------+------------+-----------------------------------------------+--------------+-------------+------------------+-----------+
| ID | TYPE |    IP     | PORT | RWSTATUS  | DISKSTATE  | COPYSETNUM |                  MOUNTPOINT                   | DISKCAPACITY |  DISKUSED   | UNHEALTHYCOPYSET |  EXTADDR  |
+----+------+-----------+------+-----------+------------+------------+-----------------------------------------------+--------------+-------------+------------------+-----------+
| 1  | nvme | 127.0.0.1 | 8201 | READWRITE | DISKNORMAL | 100        | local:///curvebs/playground/chunkserver1/data | 39 GiB       | 42140479488 | 0 %              | 127.0.0.1 |
+----+      +           +------+           +            +------------+-----------------------------------------------+--------------+-------------+------------------+           +
| 2  |      |           | 8202 |           |            | 100        | local:///curvebs/playground/chunkserver2/data | 39 GiB       | 42140479488 | 0 %              |           |
+----+      +           +------+           +            +------------+-----------------------------------------------+--------------+-------------+------------------+           +
| 3  |      |           | 8200 |           |            | 100        | local:///curvebs/playground/chunkserver0/data | 39 GiB       | 42140479488 | 0 %              |           |
+----+------+-----------+------+-----------+------------+------------+-----------------------------------------------+--------------+-------------+------------------+-----------+
```

###### list scan-status

list curvebs all copyset that scanning is false

```bash
curve bs list scan-status
```

Output:

```bash
+-------------+-----------+
| LOGICALPOOL | COPYSETID |
+-------------+-----------+
| 1           | 1         |
+-------------+-----------+
| 1           | 10        |
+-------------+-----------+
| 1           | 100       |
+-------------+-----------+
| 1           | 11        |
+-------------+-----------+
| 1           | 12        |
+-------------+-----------+
| 1           | 13        |
+-------------+-----------+
| 1           | 14        |
+-------------+-----------+
| 1           | 15        |
+-------------+-----------+
| 1           | 16        |
+-------------+-----------+
| 1           | 17        |
+-------------+-----------+
| 1           | 18        |
+-------------+-----------+
| 1           | 19        |
+-------------+-----------+
| 1           | 2         |
+-------------+-----------+
| 1           | 20        |
+-------------+-----------+
| 1           | 21        |
+-------------+-----------+
| 1           | 22        |
+-------------+-----------+
| 1           | 23        |
+-------------+-----------+
| 1           | 24        |
+-------------+-----------+
| 1           | 25        |
+-------------+-----------+
| 1           | 26        |
+-------------+-----------+
| 1           | 27        |
+-------------+-----------+
| 1           | 28        |
+-------------+-----------+
| 1           | 29        |
+-------------+-----------+
| 1           | 3         |
+-------------+-----------+
| 1           | 30        |
+-------------+-----------+
| 1           | 31        |
+-------------+-----------+
| 1           | 32        |
+-------------+-----------+
| 1           | 33        |
+-------------+-----------+
| 1           | 34        |
+-------------+-----------+
| 1           | 35        |
+-------------+-----------+
| 1           | 36        |
+-------------+-----------+
| 1           | 37        |
+-------------+-----------+
| 1           | 38        |
+-------------+-----------+
| 1           | 39        |
+-------------+-----------+
| 1           | 4         |
+-------------+-----------+
| 1           | 40        |
+-------------+-----------+
| 1           | 41        |
+-------------+-----------+
| 1           | 42        |
+-------------+-----------+
| 1           | 43        |
+-------------+-----------+
| 1           | 44        |
+-------------+-----------+
| 1           | 45        |
+-------------+-----------+
| 1           | 46        |
+-------------+-----------+
| 1           | 47        |
+-------------+-----------+
| 1           | 48        |
+-------------+-----------+
| 1           | 49        |
+-------------+-----------+
| 1           | 5         |
+-------------+-----------+
| 1           | 50        |
+-------------+-----------+
| 1           | 51        |
+-------------+-----------+
| 1           | 52        |
+-------------+-----------+
| 1           | 53        |
+-------------+-----------+
| 1           | 54        |
+-------------+-----------+
| 1           | 55        |
+-------------+-----------+
| 1           | 56        |
+-------------+-----------+
| 1           | 57        |
+-------------+-----------+
| 1           | 58        |
+-------------+-----------+
| 1           | 59        |
+-------------+-----------+
| 1           | 6         |
+-------------+-----------+
| 1           | 60        |
+-------------+-----------+
| 1           | 61        |
+-------------+-----------+
| 1           | 62        |
+-------------+-----------+
| 1           | 63        |
+-------------+-----------+
| 1           | 64        |
+-------------+-----------+
| 1           | 65        |
+-------------+-----------+
| 1           | 66        |
+-------------+-----------+
| 1           | 67        |
+-------------+-----------+
| 1           | 68        |
+-------------+-----------+
| 1           | 69        |
+-------------+-----------+
| 1           | 7         |
+-------------+-----------+
| 1           | 70        |
+-------------+-----------+
| 1           | 71        |
+-------------+-----------+
| 1           | 72        |
+-------------+-----------+
| 1           | 73        |
+-------------+-----------+
| 1           | 74        |
+-------------+-----------+
| 1           | 75        |
+-------------+-----------+
| 1           | 76        |
+-------------+-----------+
| 1           | 77        |
+-------------+-----------+
| 1           | 78        |
+-------------+-----------+
| 1           | 79        |
+-------------+-----------+
| 1           | 8         |
+-------------+-----------+
| 1           | 80        |
+-------------+-----------+
| 1           | 81        |
+-------------+-----------+
| 1           | 82        |
+-------------+-----------+
| 1           | 83        |
+-------------+-----------+
| 1           | 84        |
+-------------+-----------+
| 1           | 85        |
+-------------+-----------+
| 1           | 86        |
+-------------+-----------+
| 1           | 87        |
+-------------+-----------+
| 1           | 88        |
+-------------+-----------+
| 1           | 89        |
+-------------+-----------+
| 1           | 9         |
+-------------+-----------+
| 1           | 90        |
+-------------+-----------+
| 1           | 91        |
+-------------+-----------+
| 1           | 92        |
+-------------+-----------+
| 1           | 93        |
+-------------+-----------+
| 1           | 94        |
+-------------+-----------+
| 1           | 95        |
+-------------+-----------+
| 1           | 96        |
+-------------+-----------+
| 1           | 97        |
+-------------+-----------+
| 1           | 98        |
+-------------+-----------+
| 1           | 99        |
+-------------+-----------+
```

###### list may-broken-vol

list may broken volumes

Usage:

```bash
curve bs list may-broken-vol
```

Output:

```bash
+----------+
| FILENAME | 
+----------+
|   test   |        
+----------+
```

###### list snapshot

list curvebs snapshot

Usage:

```bash
curve bs list snapshot
```

Output:

```bash
+--------------------------------------+--------------+------+--------+----------------+-------------+----------+---------------------+---------------+
|              SNAPSHOTID              | SNAPSHOTNAME | USER | STATUS | SNAPSHOTSEQNUM | FILELENGTH  | PROGRESS |     CREATETIME      |     FILE      |
+--------------------------------------+--------------+------+--------+----------------+-------------+----------+---------------------+---------------+
| 807fdac2-5b47-42dc-b884-a4f33b0f2a1a | testsnap     | root | 0      | 1              | 10737418240 | 100      | 2023-10-15 16:16:31 | /test/test111 |
+--------------------------------------+--------------+------+--------+----------------+-------------+----------+---------------------+               +
| 204a5316-99cd-44b7-bc58-be6a547b8469 | testsnap1    | root | 0      | 2              | 10737418240 | 100      | 2023-10-25 09:22:12 |               |
```

#### clean-recycle

clean the recycle bin 

Usage:

```bash
curve bs clean-recycle --recycleprefix=/test --expiredtime=1h
```

Output:

```bash
+---------+
| RESULT  | 
+---------+
| success |        
+---------+
```

#### query

###### query file

query the file info and actual space

Usage:

```bash
curve bs query file --path=/test
```

Output:

```bash
+------+------+----------------+-------+--------+---------+--------+-----+---------------------+--------------+---------+-----------------+----------+
|  ID  | NAME |      TYPE      | OWNER | CHUNK  | SEGMENT | LENGTH | SEQ |        CTIME        |    STATUS    | STRIPE  |    THROTTLE     |  ALLOC   |
+------+------+----------------+-------+--------+---------+--------+-----+---------------------+--------------+---------+-----------------+----------+
| 1003 | test | INODE_PAGEFILE | test  | 16 MiB | 1.0 GiB | 10 GiB | 1   | 2022-08-29 17:00:55 | kFileCreated | count:0 | type:IOPS_TOTAL | size:0 B |
|      |      |                |       |        |         |        |     |                     |              | uint:0  | limit:2000      |          |
|      |      |                |       |        |         |        |     |                     |              |         | type:BPS_TOTAL  |          |
|      |      |                |       |        |         |        |     |                     |              |         | limit:125829120 |          |
+------+------+----------------+-------+--------+---------+--------+-----+---------------------+--------------+---------+-----------------+----------+
```

###### query chunk

query the location of the chunk corresponding to the offset

Usage:

```bash
curve bs query chunk --path /test1 --offset 1008600000 
```

Output:

```bash
+-------+-------------+---------+------------+----------------------+
| CHUNK | LOGICALPOOL | COPYSET |   GROUP    |       LOCATION       |
+-------+-------------+---------+------------+----------------------+
| 61    | 1           | 61      | 4294967357 | ***.***.***.***:**** |
|       |             |         |            | ***.***.***.***:**** |
|       |             |         |            | ***.***.***.***:**** |
+-------+-------------+---------+------------+----------------------+
```

###### query segment

query the segments info of the file

Usage:

```bash
curve bs query seginfo --path /test1 
```

Output:

```bash
+-------------+-------------+-----------+------------+---------+-------+
| LOGICALPOOL | SEGMENTSIZE | CHUNKSIZE |   START    | COPYSET | CHUNK |
+-------------+-------------+-----------+------------+---------+-------+
| 1           | 1073741824  | 16777216  | 0          | 1       | 1     |
+             +             +           +            +---------+-------+
|                        ......                                        |
+             +             +           +------------+---------+-------+
|             |             |           | 9663676416 | 1       | 101   |
+             +             +           +            +---------+-------+
|                        ......                                        |
+             +             +           +            +---------+-------+
|             |             |           |            | 99      | 99    |
+-------------+-------------+-----------+------------+---------+-------+
```

##### query scan-status
quey ScanStatus Info in bs

Usage:
```bash
curve bs query scan-satus --copysetid 1 --logicalpoolid 1
```

Output:
```bash
+-------------+-----------+-------+-------------+--------------------+
| LOGICALPOOL | COPYSETID | SCAN  | LASTSCANSEC | LASTSCANCONSISTENT |
+-------------+-----------+-------+-------------+--------------------+
| 1           | 1         | false | 1684425801  | true               |
+-------------+-----------+-------+-------------+--------------------+
```

#### status

##### status etcd
get the etcd status of curvebs

Usage:

```bash
curve bs status etcd
```

Output:

```bash
+---------------------+---------+----------+
|        ADDR         | VERSION |  STATUS  |
+---------------------+---------+----------+
| ***.***.*.***:***** | 3.4.10  | follower |
+---------------------+         +          +
| ***.***.*.***:***** |         |          |
+---------------------+         +----------+
| ***.***.*.***:***** |         | leader   |
+---------------------+---------+----------+
```

##### status mds

get the mds status of curvebs

Usage:

```bash
curve bs status mds
```

Output:

```bash
+-------------------+-------------------+-------------------+----------+
|       ADDR        |     DUMMYADDR     |      VERSION      |  STATUS  |
+-------------------+-------------------+-------------------+----------+
| **.***.**.**:**** | **.***.**.**:**** | ci+562296c7+debug | follower |
+-------------------+-------------------+                   +          +
| **.***.**.**:**** | **.***.**.**:**** |                   |          |
+-------------------+-------------------+                   +----------+
| **.***.**.**:**** | **.***.**.**:**** |                   | leader   |
+-------------------+-------------------+-------------------+----------+
```

##### status client

get the client status of curvebs

Usage:

```bash
curve bs status client
```

Output:

```bash
+-------------+----------------+---------------------+-----+
|    TYPE     |    VERSION     |        ADDR         | NUM |
+-------------+----------------+---------------------+-----+
| nebd-server | 9.9.9+2c4861ca | ***.***.**.***:**** | 2   |
+             +                +---------------------+     +
|             |                | ***.***.**.***:**** |     |
+-------------+----------------+---------------------+-----+
```

##### status snapshotserver

get the mds status of curvebs

Usage:

```bash
curve bs status snapshotserver
```

Output:

```bash
+---------------------+---------------------+-------------------+----------+
|        ADDR         |      DUMMYADDR      |    VERSION        |  STATUS  |
+---------------------+---------------------+-------------------+----------+
| ***.***.**.***:**** | ***.***.**.***:**** | ci+562296c7+debug | follower |
+---------------------+---------------------+                   +          +
| ***.***.**.***:**** | ***.***.**.***:**** |                   |          |
+---------------------+---------------------+                   +----------+
| ***.***.**.***:**** | ***.***.**.***:**** |                   | leader   |
+---------------------+---------------------+-------------------+----------+
```

##### status chunkserver

get the chunkserver status of curvebs

Usage:

```bash
curve bs status chunkserver
```

Output:

```bash
+------------------+------------------+----------------+--------+------------+
|   EXTERNALADDR   |   INTERNALADDR   |    VERSION     | STATUS | RECOVERING |
+------------------+------------------+----------------+--------+------------+
| **************** | **************** | d9b6bb98+debug | online | false      |
+------------------+------------------+                +        +            +
| **************** | **************** |                |        |            |
+------------------+------------------+                +        +            +
| **************** | **************** |                |        |            |
+------------------+------------------+----------------+--------+------------+
```

##### status copyset

get the copyset status of curvebs

Usage:

```bash
curve bs status copyset
```

Output:

```bash
+------------+-----------+--------+--------+--------+---------+
| COPYSETKEY | COPYSETID | POOLID | STATUS | LOGGAP | EXPLAIN |
+------------+-----------+--------+--------+--------+---------+
| 4294967297 | 1         | 1      | ok     | 0      |         |
+------------+-----------+        +        +--------+---------+
| ......     |  ......   | ...... | ...... | ...... | ......  |
+------------+-----------+        +        +--------+---------+
| 4294967395 | 99        |        |        | 0      |         |
+------------+-----------+--------+--------+--------+---------+
```

#### delete

##### delete peer

delete the peer from the copyset

Usage:

```bash
curve bs delete peer
```

Output:

```bash
+------------------+------------------+---------+---------+--------+
|      LEADER      |       PEER       | COPYSET | RESULT  | REASON |
+------------------+------------------+---------+---------+--------+
| 127.0.0.1:8201:0 | 127.0.0.1:8202:0 | (1:29)  | success | null   |
+------------------+------------------+---------+---------+--------+
```

##### delete volume

###### delete volume clone

delete volume clone tasks in curvebs cluster

Usage:

```bash
curve bs delete volume clone
```

Output:

```bash
+------+--------------------------------------+--------------------------------------+-------+---------+            
| USER |                 SRC                  |               TASK ID                | FILE  | RESULT  |            
+------+--------------------------------------+--------------------------------------+-------+---------+            
| root | a19b5e5e-b306-488f-8e6d-d87282c869cb | d26e27a8-fcbd-4f7a-adf8-53795217cbb0 | /root | success |            
+------+--------------------------------------+--------------------------------------+-------+---------+  
```

###### delete volume recover

delete volume recover tasks in curvebs cluster

Usage:

```bash
curve bs delete volume recover
```

Output:

```bash
+------+--------------------------------------+--------------------------------------+-------+---------+             
| USER |                 SRC                  |               TASK ID                | FILE  | RESULT  |             
+------+--------------------------------------+--------------------------------------+-------+---------+             
| root | a19b5e5e-b306-488f-8e6d-d87282c869cb | 9dfa8699-a275-4891-8ec2-e447a0ccc77c | /root | success |             
+------+--------------------------------------+--------------------------------------+-------+---------+   
```

#### update

##### update peer

reset peer

Usage:
```bash
curve bs update peer 127.0.0.0:8200:0 --logicalpoolid=1 --copysetid=1
```

Output:
```
+----------------------+---------+---------+--------+
|         PEER         | COPYSET | RESULT  | REASON |
+----------------------+---------+---------+--------+
|   127.0.0.0:8200:0   | (1:1)   | success | null   |
+----------------------+---------+---------+--------+
```

##### update leader

transfer leader

Usage:
```bash
curve bs update leader 127.0.0.1:8202:0 --logicalpoolid=1 --copysetid=1 --peers=127.0.0.1:8200:0,127.0.0.1:8201:0,127.0.0.1:8202:0
```

Output:
```
+-----------------------+-----------------------+---------+---------+
|        LEADER         |       OLDLEADER       | COPYSET | RESULT  |
+-----------------------+-----------------------+---------+---------+
| ***.***.**.***:****:* | ***.***.**.***:****:* | (1:1)   | success |
+-----------------------+-----------------------+---------+---------+
```

##### update file

expand pagefile

Usage:
```bash
curve bs update file --path /test2/test1 --size 10
```

Output:
```
+---------+
| RESULT  |
+---------+
| success |
+---------+
```

##### update throttle

update file throttle params

Usage:
```bash
curve bs update throttle --path /test1 --type=bps_total --limit 20000
```

Output:
```
+---------+
| RESULT  |
+---------+
| success |
+---------+
```

##### update scan-state

enable/disable scan for logical pool

Usage:
```bash
curve bs update scan-state --logicalpoolid 1 [--scan=true/false]
```

Output:
```
+----+------+---------+--------+
| ID | SCAN | RESULT  | REASON |
+----+------+---------+--------+
| 1  | true | success | null   |
+----+------+---------+--------+
```

##### update copyset availflag

update copyset availflag

Usage:
```bash
curve bs update copyset availflag --availflag=true [--dryrun=true/false]
```

Output:
```
+--------+-----------+---------------+--------+
| POOLID | COPYSETID |   AVAILFLAG   | DRYRUN |
+--------+-----------+---------------+--------+
| 1      | 1         | false => true | true   |
+--------+-----------+---------------+--------+
```

##### update leader-schedule

"rapidly transfer leader

Usage:
```bash
curve bs update leader-schedule --logicalpoolid 1
curve bs update leader-schedule --all
```

Output:
```
+---------+--------+
| RESULT  | REASON |
+---------+--------+
| success | null   |
+---------+--------+
```

##### update volume flatten

update volume flatten in curvebs cluster

Usage:
```bash
curve bs update volume --user root --taskid d26e27a8-fcbd-4f7a-adf8-53795217cbb0
```

Output:
```
+------+--------------------------------------+---------+            
| USER |               TASK ID                | RESULT  |            
+------+--------------------------------------+---------+            
| root | d26e27a8-fcbd-4f7a-adf8-53795217cbb0 | success |            
+------+--------------------------------------+---------+  
```

#### create

##### create file

create pagefile

Usage:
```bash
curve bs create file --path /test2/test4  --size 10
```

Output:
```
+---------+
| RESULT  |
+---------+
| success |
+---------+
```

##### create dir

create directory

Usage:
```bash
curve bs create dir --path /test2/test5 
```

Output:
```
+---------+
| RESULT  |
+---------+
| success |
+---------+
```

#### check

##### check copyset

check copysets health in curvebs

Usage:

```shell
curve bs check copyset --copysetid 1 --logicalpoolid 1
```

Output:

```shell
+------------+-----------+--------+--------+--------+---------+
| COPYSETKEY | COPYSETID | POOLID | STATUS | LOGGAP | EXPLAIN |
+------------+-----------+--------+--------+--------+---------+
| 4294967297 | 1         | 1      | ok     | 0      |         |
+------------+-----------+--------+--------+--------+---------+
```
##### check chunkserver

check chunkserver health in curvebs

Usage:

```shell
curve bs check chunkserver --chunkserverid 1
```

Output:

```shell
+------------+-----------+--------+--------+--------+---------+
| CHUNKSERVERID | HELATHYCOUNT | UNHEALTHYCOUNT | UNHEALTHYRATIO |
+------------+-----------+--------+--------+--------+---------+
| 1             | 100          | 0              | 0.00%          |   
+------------+-----------+--------+--------+--------+---------+
```

##### check server

check copysets health in server

Usage:

```shell
curve bs check server --serverid 1
curve bs check server --ip 127.0.0.1 --port 8200
```

Output:

```shell
+--------+-----------+-------+------------------+
| SERVER |    IP     | TOTAL | UNHEALTHYCOPYSET |
+--------+-----------+-------+------------------+
| 1      | 127.0.0.1 | 100   | 0(0%)            |
+--------+-----------+-------+------------------+
```

##### check consistency

check the file consistency

Usage:

```shell
curve bs check consistency --path /test
```

Output:

```shell
+--------+--------+------------+
|  NAME  | STATUS |   EXPLAIN  |
+--------+--------+------------+
| /test  |  ok    |      -     |
+--------+--------+------------+
```

#### snapshot

##### snapshot copyset

take snapshot for copyset

Usage:
```bash
curve bs snapshot copyset 127.0.0.0:8200:0 --logicalpoolid=1 --copysetid=1
```

Output:
```
+-----------------------+---------+---------+
|         PEER          | COPYSET | RESULT  |
+-----------------------+---------+---------+
| ***.***.**.***:****:* | (**:**) | success |
+-----------------------+---------+---------+
```

```bash
curve bs snapshot copyset --all
```

Output:
```
+----------------+---------+
|  CHUNKSERVER   | RESULT  |
+----------------+---------+
| **.*.*.**:8200 | failed  |
+----------------+---------+
| **.*.*.**:8201 | success |
+----------------+         +
| **.*.*.**:8202 |         |
+----------------+---------+
```

## Comparison of old and new commands

### curve fs

| old                            | new                        |
| ------------------------------ | -------------------------- |
| curvefs_tool check-copyset     | curve fs check copyset     |
| curvefs_tool create-fs         | curve fs create fs         |
| curvefs_tool create-topology   | curve fs create topology   |
| curvefs_tool delete-fs         | curve fs delete fs         |
| curvefs_tool list-copyset      | curve fs list copyset      |
| curvefs_tool list-fs           | curve fs list fs           |
| curvefs_tool list-fs           | curve fs list mountpoint   |
| curvefs_tool list-partition    | curve fs list partition    |
| curvefs_tool query-copyset     | curve fs query copyset     |
| curvefs_tool query-fs          | curve fs query fs          |
| curvefs_tool query-inode       | curve fs query inode       |
| curvefs_tool query-metaserver  | curve fs query metaserver  |
| curvefs_tool query-partition   | curve fs query partition   |
| curvefs_tool status-mds        | curve fs status mds        |
| curvefs_tool status-metaserver | curve fs status metaserver |
| curvefs_tool status-etcd       | curve fs status etcd       |
| curvefs_tool status-copyset    | curve fs status copyset    |
| curvefs_tool status-cluster    | curve fs status cluster    |
| curvefs_tool umount-fs         | curve fs umount fs         |
| curvefs_tool usage-inode       | curve fs usage inode       |
| curvefs_tool usage-metadata    | curve fs usage metadata    |

### curve bs

| old                                  | new                               |
| ------------------------------------ | --------------------------------- |
| curve_ops_tool logical-pool-list     | curve bs list logical-pool        |
| curve_ops_tool get -fileName=        | curve bs query file -path         |
| curve_ops_tool etcd-status           | curve bs status etcd              |
| curve_ops_tool mds-status            | curve bs status mds               |
| curve_ops_tool server-list           | curve bs list server              |
| curve_ops_tool client-list           | curve bs list client              |
| curve_ops_tool delete                | curve bs delete file              |
| curve_ops_tool list                  | curve bs list dir                 |
| curve_ops_tool create                | curve bs create file/dir          |
| curve_ops_tool seginfo               | curve bs query seginfo            |
| curve_ops_tool chunk-location        | curve bs query chunk              |
| curve_ops_tool remove-peer           | curve bs delete peer              |
| curve_ops_tool reset-peer            | curve bs update peer              |
| curve_ops_tool space                 | curve bs list space               |
| curve_ops_tool update-throttle       | curve bs update throttle          |
| curve_ops_tool check-copyset         | curve bs check copyset            |
| curve_ops_tool client-status         | curve bs status client            |
| curve_ops_tool check-operator        | curve bs check operator           |
| curve_ops_tool snapshot-clone-status | curve bs status snapshotserver    |
| curve_ops_tool transfer-leader       | curve bs update leader            |
| curve_ops_tool do-snapshot           | curve bs snapshot copyset         |
| curve_ops_tool set-scan-state        | curve bs update scan-state        |
| curve_ops_tool chunkserver-status    | curve bs status chunkserver       |
| curve_ops_tool chunkserver-list      | curve bs list chunkserver         |
| curve_ops_tool set-copyset-availflag | curve bs update copyset availflag |
| curve_ops_tool scan-status           | curve bs list/query scan-status   |
| curve_ops_tool clean-recycle         | curve bs clean-recycle            |
| curve_ops_tool copysets-status       | curve bs status copyset           |
| curve_ops_tool list-may-broken-vol   | curve bs list may-broken-vol      |
| curve_ops_tool rapid-leader-schedule | curve bs update leader-schedule   |
| curve_ops_tool do-snapshot-all       | curve bs snapshot --all           |
| curve_ops_tool check-chunkserver     | curve bs check chunkserver        |
| curve_ops_tool status                |                                   |
| curve_ops_tool check-consistency     | curve bs check consistency        |
| curve_ops_tool check-server          | curve bs check server             |

