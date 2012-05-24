package account

import (
    "payment_center/mysql"
    "time"
)

type TransferLog struct {
    Id             int64
    TransferId     int64
    PayerAccountId int64
    PayeeAccountId int64
    Amount         int64
    StatusCode     byte
    Status         string
    Memo           string
    CreateTime     int64
}

func (tl *TransferLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "transfer_log SET " +
        "transfer_id=?, " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "amount =?, " +
        "status_code =?, " +
        "status=?, " +
        "memo=?, " +
        "create_time=?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    tl.CreateTime = now
    res, err := ins.Exec(
        tl.TransferId,
        tl.PayerAccountId,
        tl.PayeeAccountId,
        tl.Amount,
        tl.StatusCode,
        tl.Status,
        tl.Memo,
        tl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
