package account

import (
    "payment_center/mysql"
    "strconv"
    "time"
)

type Account struct {
    Id                int64
    Password          string
    Freeze            byte
    Currency          byte
    TotalAmount       int64
    UseAmount         int64
    WithdrawalsAmount int64
    FreezeAmount      int64
    UnSettledAmount   int64
    CreateTime        int64
    UpdateTime        int64
}

func (a *Account) Create() error {
    if a.Id > 0 {
        return setError("已经存在此账户")
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "account SET " +
        "password=?, " +
        "freeze=?, " +
        "currency=?, " +
        "total_amount=?, " +
        "use_amount=?, " +
        "withdrawals_amount=?, " +
        "freeze_amount=?, " +
        "unsettled_amount=?, " +
        "create_time=?, " +
        "update_time=?")
    if err != nil {
        return err
    }

    now := time.Now().Unix()
    a.CreateTime = now
    a.UpdateTime = now
    res, err := ins.Exec(
        a.Password,
        a.Freeze,
        a.Currency,
        a.TotalAmount,
        a.UseAmount,
        a.WithdrawalsAmount,
        a.FreezeAmount,
        a.UnSettledAmount,
        a.CreateTime,
        a.UpdateTime,
    )
    if err != nil {
        return err
    }

    aid, err := res.LastInsertId()
    a.Id = aid
    al := &AccountLog{
        AccountId:     a.Id,
        OptTypeCode:   1,
        OptType:       "创建",
        OptId:         a.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "创建账号，编号：" + strconv.FormatInt(aid, 10),
    }

    _, err = al.Create()
    if err != nil {
        return err
    }

    return nil
}

func (a *Account) Update() (int64, error) {
    if a.Id <= 0 {
        return 0, setError("无此账户")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "account SET " +
        "password=?, " +
        "freeze=?, " +
        "currency=?, " +
        "total_amount=?, " +
        "use_amount=?, " +
        "withdrawals_amount=?, " +
        "freeze_amount=?, " +
        "unsettled_amount=?, " +
        "update_time=? " +
        "WHERE id = ?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    a.UpdateTime = now
    res, err := upd.Exec(
        a.Password,
        a.Freeze,
        a.Currency,
        a.TotalAmount,
        a.UseAmount,
        a.WithdrawalsAmount,
        a.FreezeAmount,
        a.UnSettledAmount,
        a.UpdateTime,
        a.Id,
    )
    if err != nil {
        return 0, err
    }

    rowsAN, err := res.RowsAffected()
    al := &AccountLog{
        AccountId:     a.Id,
        OptTypeCode:   2,
        OptType:       "更新",
        OptId:         a.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "更新账号，编号：" + strconv.FormatInt(a.Id, 10),
    }
    _, err = al.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}

func (a *Account) Get() error {
    if a.Id <= 0 {
        return setError("无此账户")
    }

    row := doDb.QueryRow("SELECT * FROM "+mysql.PreTable+"account WHERE id =? LIMIT 1", a.Id)
    err := row.Scan(
        &a.Id,
        &a.Password,
        &a.Freeze,
        &a.Currency,
        &a.TotalAmount,
        &a.UseAmount,
        &a.WithdrawalsAmount,
        &a.FreezeAmount,
        &a.UnSettledAmount,
        &a.CreateTime,
        &a.UpdateTime,
    )
    if err != nil {
        return setError("无此账户")
    }

    return nil
}

func (a *Account) GetNoFreeze() error {
    err := a.Get()
    if err != nil {
        return err
    }

    if a.Freeze != 0 {
        return setError("此账户被冻结")
    }

    return nil
}
