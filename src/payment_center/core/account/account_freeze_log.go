package account

import (
    "payment_center/mysql"
    "time"
)

type AccountFreezeLog struct {
    Id              int64
    AccountFreezeId int64
    AccountId       int64
    TypeCode        byte
    Type            string
    StatusCode      byte
    Status          string
    Memo            string
    CreateTime      int64
}

func (afl *AccountFreezeLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "account_freeze_log SET " +
        "account_freeze_id=?, " +
        "account_id=?, " +
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
    afl.CreateTime = now
    res, err := ins.Exec(
        afl.AccountFreezeId,
        afl.AccountId,
        afl.TypeCode,
        afl.Type,
        afl.StatusCode,
        afl.Status,
        afl.Memo,
        afl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
