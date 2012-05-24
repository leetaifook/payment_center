package account

import (
    "payment_center/mysql"
    "time"
)

type AccountLog struct {
    Id                int64
    AccountId         int64
    OptTypeCode       byte
    OptType           string
    OptId             int64
    OptStatusCode     byte
    OptStatus         string
    FreezeCode        byte
    Freeze            string
    CurrencyCode      byte
    Currency          string
    TotalAmount       int64
    UseAmount         int64
    WithdrawalsAmount int64
    FreezeAmount      int64
    UnSettledAmount   int64
    Memo              string
    CreateTime        int64
}

func (al *AccountLog) Create() (int64, error) {
    a := &Account{Id: al.AccountId}
    a.Get()

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "account_log SET " +
        "account_id=?, " +
        "opt_type_code=?, " +
        "opt_type =?, " +
        "opt_id=?, " +
        "opt_status_code=?, " +
        "opt_status=?, " +
        "freeze_code=?, " +
        "freeze=?, " +
        "currency_code=?, " +
        "currency=?, " +
        "total_amount=?, " +
        "use_amount=?, " +
        "withdrawals_amount=?, " +
        "freeze_amount=?, " +
        "unsettled_amount=?, " +
        "memo=?, " +
        "create_time=?")
    if err != nil {
        return 0, err
    }

    al.Currency = "人民币"
    if a.Freeze == 0 {
        al.Freeze = "解冻"
    } else {
        al.Freeze = "冻结"
    }

    now := time.Now().Unix()
    al.CreateTime = now
    res, err := ins.Exec(
        a.Id,
        al.OptTypeCode,
        al.OptType,
        al.OptId,
        al.OptStatusCode,
        al.OptStatus,
        a.Freeze,
        al.Freeze,
        a.Currency,
        al.Currency,
        a.TotalAmount,
        a.UseAmount,
        a.WithdrawalsAmount,
        a.FreezeAmount,
        a.UnSettledAmount,
        al.Memo,
        al.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
