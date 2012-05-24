package account

import (
    "payment_center/mysql"
    "time"
)

type WithdrawalsLog struct {
    Id            int64
    WithdrawalsId int64
    AccountId     int64
    Amount        int64
    StatusCode    byte
    Status        string
    Memo          string
    CreateTime    int64
}

func (wl *WithdrawalsLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "withdrawals_log SET " +
        "withdrawals_id=?, " +
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
    wl.CreateTime = now
    res, err := ins.Exec(
        wl.WithdrawalsId,
        wl.AccountId,
        wl.Amount,
        wl.StatusCode,
        wl.Status,
        wl.Memo,
        wl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
