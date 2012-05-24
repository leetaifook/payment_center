package account

import (
    "database/sql"
    "payment_center/mysql"
)

var (
    transaction bool = false
    doDb        AccountDb
    db          *sql.DB = mysql.Db
    tdb         *sql.Tx
)

type AccountDb interface {
    Exec(query string, args ...interface{}) (sql.Result, error)
    Prepare(query string) (*sql.Stmt, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
}
