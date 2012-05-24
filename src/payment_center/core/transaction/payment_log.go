package transaction

import (
    "payment_center/mysql"
    "time"
)

type PaymentLog struct {
    Id           int64
    PaymentId    int64
    Amount       int64
    TypeCode     byte
    Type         string
    StatusCode   byte
    Status       string
    CurrencyCode byte
    Currency     string
    Memo         string
    CreateTime   int64
}

func (pl *PaymentLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "payment_log SET " +
        "payment_id=?, " +
        "amount=?, " +
        "type_code=?, " +
        "type=?, " +
        "status_code=?, " +
        "status=?, " +
        "currency_code=?, " +
        "currency=?, " +
        "memo=?, " +
        "create_time=?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    pl.CreateTime = now
    res, err := ins.Exec(
        pl.PaymentId,
        pl.Amount,
        pl.TypeCode,
        pl.Type,
        pl.StatusCode,
        pl.Status,
        pl.CurrencyCode,
        pl.Currency,
        pl.Memo,
        pl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
