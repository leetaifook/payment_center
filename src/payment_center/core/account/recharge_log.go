package account

import (
    "payment_center/mysql"
    "time"
)

type RechargeLog struct {
    Id         int64
    RechargeId int64
    AccountId  int64
    Amount     int64
    StatusCode byte
    Status     string
    Memo       string
    CreateTime int64
}

func (rl *RechargeLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "recharge_log SET " +
        "recharge_id=?, " +
        "account_id=?, " +
        "amount=?, " +
        "status_code=?, " +
        "status=?, " +
        "memo=?, " +
        "create_time=?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    rl.CreateTime = now
    res, err := ins.Exec(
        rl.RechargeId,
        rl.AccountId,
        rl.Amount,
        rl.StatusCode,
        rl.Status,
        rl.Memo,
        rl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
