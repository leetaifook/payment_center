package account

import (
    "payment_center/mysql"
    "time"
)

type FundsFreezeLog struct {
    Id            int64
    FundsFreezeId int64
    AccountId     int64
    Amount        int64
    TypeCode      byte
    Type          string
    StatusCode    byte
    Status        string
    Memo          string
    CreateTime    int64
}

func (ffl *FundsFreezeLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "funds_freeze_log SET " +
        "funds_freeze_id=?, " +
        "account_id=?, " +
        "amount=?, " +
        "type_code =?, " +
        "type=?, " +
        "status_code =?, " +
        "status=?, " +
        "memo=?, " +
        "create_time=?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    ffl.CreateTime = now
    res, err := ins.Exec(
        ffl.FundsFreezeId,
        ffl.AccountId,
        ffl.Amount,
        ffl.TypeCode,
        ffl.Type,
        ffl.StatusCode,
        ffl.Status,
        ffl.Memo,
        ffl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
