package transaction

import (
    "payment_center/mysql"
    "time"
)

type ReceivablesLog struct {
    Id            int64
    ReceivablesId int64
    Amount        int64
    TypeCode      byte
    Type          string
    StatusCode    byte
    Status        string
    CurrencyCode  byte
    Currency      string
    Memo          string
    CreateTime    int64
}

func (rl *ReceivablesLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "receivables_log SET " +
        "receivables_id=?, " +
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
    rl.CreateTime = now
    res, err := ins.Exec(
        rl.ReceivablesId,
        rl.Amount,
        rl.TypeCode,
        rl.Type,
        rl.StatusCode,
        rl.Status,
        rl.CurrencyCode,
        rl.Currency,
        rl.Memo,
        rl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
