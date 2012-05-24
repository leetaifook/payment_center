package transaction

import (
    "database/sql"
    "payment_center/mysql"
)

var (
    transaction bool = false
    doDb        TransactionDb
    db          *sql.DB = mysql.Db
    tdb         *sql.Tx
)

type TransactionDb interface {
    Exec(query string, args ...interface{}) (sql.Result, error)
    Prepare(query string) (*sql.Stmt, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
}
