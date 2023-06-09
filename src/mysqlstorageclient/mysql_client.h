#ifndef SRC_MYSQLSTORAGECLIENT_MYSQL_CLIENT_H_
#define SRC_MYSQLSTORAGECLIENT_MYSQL_CLIENT_H_

#include <string>
#include <vector>
#include <utility>
#include <mysql_connection.h>
#include <mysql_driver.h>
#include <cppconn/driver.h>
#include <cppconn/exception.h>
#include <cppconn/prepared_statement.h>
#include <cppconn/statement.h>
#include <cppconn/resultset.h>

namespace curve {
namespace mysqlstorage {


enum class MySQLError {
    MysqlOK= 0,
    ConnectionError = 2000,
    StatementError = 2001,
    QueryError = 2002,
    // 添加其他 MySQL 错误代码
};

class MysqlConf {
 public:
    MysqlConf(std::string host = "tcp://101.201.83.65:3306",
              std::string user = "curve", std::string passwd = "021013Jbk.",
              std::string db = "curve")
        : host_(host), user_(user), passwd_(passwd), db_(db) {}
    ~MysqlConf() {}
    std::string host_;
    std::string user_;
    std::string passwd_;
    std::string db_;
};

class MysqlClientImp {
public:
    MysqlClientImp() {}
    virtual ~MysqlClientImp() {}

    int Init(MysqlConf conf, int timeout, int retryTiems);

    int CreateTable(const std::string &tableName);

    int DropTable(const std::string &tableName);

    void CloseClient();

    int Put(const std::string &key, const std::string &value) ;

    int PutRewithRevision(const std::string &key, const std::string &value,
        int64_t *revision) ;

    int Get(const std::string &key, std::string &out) ;


    int List(const std::string &startKey,
        const std::string &endKey, std::vector<std::string> *values) ;

    int List(const std::string& startKey, const std::string& endKey,
             std::vector<std::pair<std::string, std::string> >* out) ;

    //int Rename(const std::string &oldKey, const std::string &newKey) ;

    int Delete(const std::string &key) ;

    int DeleteRewithRevision(
        const std::string &key, int64_t *revision) ;

    int CompareAndSwap(const std::string &key, const std::string &preV,
        const std::string &target) ;

    int GetCurrentRevision(int64_t *revision);

    int ListWithLimitAndRevision(const std::string &startKey,
        const std::string &endKey, int64_t limit, int64_t revision,
        std::vector<std::string> *values, std::string *lastKey);


    sql::Connection *conn_;
    sql::Statement *stmt_;
    sql::PreparedStatement *prep_stmt_;
};


}
}



#endif