package transaction

import (
    "fmt"
)

func NewPayment(payerid, payeeid, amount int64) (*Payment, error) {
    begin()
    p := &Payment{
        PayerAccountId: payerid,
        PayeeAccountId: payeeid,
        Amount:         amount,
        Status:         1,
    }
    err := p.Create()
    if err != nil {
        err := rollback()
        if err != nil {
            fmt.Println(err)
        }
    } else {
        err := commit()
        if err != nil {
            fmt.Println(err)
        }
    }

    return p, err
}

func NewReceivables(payeeid, payerid, amount int64) (*Receivables, error) {
    begin()
    r := &Receivables{
        PayeeAccountId: payeeid,
        PayerAccountId: payerid,
        Amount:         amount,
        Status:         1,
    }
    err := r.Create()
    if err != nil {
        err := rollback()
        if err != nil {
            fmt.Println(err)
        }
    } else {
        err := commit()
        if err != nil {
            fmt.Println(err)
        }
    }

    return r, err
}

func setDb() {
    if transaction {
        dbTmp, err := db.Begin()
        if err != nil {
            panic(err)
        } else {
            tdb = dbTmp
        }
    }
}

func getDb() {
    if transaction {
        doDb = tdb
    } else {
        doDb = db
    }
}

func begin() {
    transaction = true
    setDb()
    getDb()
}

func commit() error {
    var err error
    if transaction {
        err = tdb.Commit()
    }

    transaction = false
    setDb()
    getDb()

    return err
}

func rollback() error {
    var err error
    if transaction {
        err = tdb.Rollback()
    }

    transaction = false
    setDb()
    getDb()

    return err
}

func setError(s string) error {
    return &TransactionError{s}
}
