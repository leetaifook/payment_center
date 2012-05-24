package account

import (
    "crypto/md5"
    "fmt"
    "io"
)

func NewAccount(password string, currency byte) (*Account, error) {
    begin()
    h := md5.New()
    io.WriteString(h, password)
    a := &Account{
        Password: fmt.Sprintf("%x", h.Sum(nil)),
        Currency: currency,
    }
    err := a.Create()
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

    return a, err
}

func NewRecharge(aid, amount int64) (*Recharge, error) {
    begin()
    r := &Recharge{
        AccountId: aid,
        Amount:    amount,
        Status:    1,
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

func NewWithdrawals(aid, amount int64) (*Withdrawals, error) {
    begin()
    w := &Withdrawals{
        AccountId: aid,
        Amount:    amount,
        Status:    1,
    }
    err := w.Create()
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

    return w, err
}

func NewTransfer(payerid, payeeid, amount int64) (*Transfer, error) {
    begin()
    t := &Transfer{
        PayerAccountId: payerid,
        PayeeAccountId: payeeid,
        Amount:         amount,
        Status:         1,
    }
    err := t.Create()
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

    return t, err
}

func NewAccountFreeze(aid int64, freezeType byte, reason string) (*AccountFreeze, error) {
    begin()
    af := &AccountFreeze{
        AccountId: aid,
        Type:      freezeType,
        Status:    1,
        Reason:    reason,
    }
    err := af.Create()
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

    return af, err
}

func NewFundsFreeze(aid, amount int64, freezeType byte, reason string) (*FundsFreeze, error) {
    begin()
    ff := &FundsFreeze{
        AccountId: aid,
        Amount:    amount,
        Type:      freezeType,
        Status:    1,
        Reason:    reason,
    }
    err := ff.Create()
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

    return ff, err
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
    return &AccountError{s}
}
