package account

import (
    "payment_center/mysql"
    "strconv"
    "time"
)

type Withdrawals struct {
    Id         int64
    AccountId  int64
    Amount     int64
    Status     byte
    Enabled    byte
    CreateTime int64
    UpdateTime int64
}

func (w *Withdrawals) Create() error {
    if w.Id > 0 {
        return setError("已经存在此提现信息")
    }

    a := &Account{Id: w.AccountId}
    err := a.GetNoFreeze()
    if err != nil {
        return err
    }

    switch {
    case a.WithdrawalsAmount <= 0:
        return setError("无可提现金额")
    case a.UseAmount < w.Amount:
        return setError("可提现金额少于提现金额")
    case a.WithdrawalsAmount < w.Amount:
        return setError("可提现金额少于提现金额")
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "withdrawals SET " +
        "account_id=?, " +
        "amount=?, " +
        "status=?, " +
        "enabled=?, " +
        "create_time=?, " +
        "update_time=?")
    if err != nil {
        return err
    }

    now := time.Now().Unix()
    w.CreateTime = now
    w.UpdateTime = now
    res, err := ins.Exec(
        w.AccountId,
        w.Amount,
        w.Status,
        w.Enabled,
        w.CreateTime,
        w.UpdateTime,
    )
    if err != nil {
        return err
    }

    wid, err := res.LastInsertId()
    w.Id = wid
    wl := &WithdrawalsLog{
        WithdrawalsId: w.Id,
        AccountId:     w.AccountId,
        Amount:        w.Amount,
        StatusCode:    1,
        Status:        "成功",
        Memo:          "新建：对账户编号：" + strconv.FormatInt(w.AccountId, 10) + "提现，提现编号：" + strconv.FormatInt(w.Id, 10) + ",提现金额为：" + strconv.FormatInt(w.Amount, 10),
    }

    _, err = wl.Create()
    if err != nil {
        return err
    } else {
        w.Enabled = 1
        _, err := w.Update()
        if err != nil {
            return err
        }
    }

    al := &AccountLog{
        AccountId:     w.AccountId,
        OptTypeCode:   4,
        OptType:       "提现",
        OptId:         w.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "提现账户编号：" + strconv.FormatInt(w.AccountId, 10),
    }
    _, err = al.Create()
    if err != nil {
        return err
    }

    a.TotalAmount -= w.Amount
    a.UseAmount -= w.Amount
    a.WithdrawalsAmount -= w.Amount
    _, err = a.Update()
    if err != nil {
        return err
    }

    return nil
}

func (w *Withdrawals) Update() (int64, error) {
    if w.Id <= 0 {
        return 0, setError("无此提现信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "withdrawals SET " +
        "account_id=?, " +
        "amount=?, " +
        "status=?, " +
        "enabled=?, " +
        "update_time=? " +
        "WHERE id = ?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    w.UpdateTime = now
    res, err := upd.Exec(
        w.AccountId,
        w.Amount,
        w.Status,
        w.Enabled,
        w.UpdateTime,
        w.Id,
    )
    if err != nil {
        return 0, err
    }

    rowsAN, err := res.RowsAffected()
    wl := &WithdrawalsLog{
        WithdrawalsId: w.Id,
        AccountId:     w.AccountId,
        Amount:        w.Amount,
        StatusCode:    1,
        Status:        "成功",
        Memo:          "更新：对账户编号：" + strconv.FormatInt(w.AccountId, 10) + "提现，提现编号：" + strconv.FormatInt(w.Id, 10) + ",提现金额为：" + strconv.FormatInt(w.Amount, 10),
    }
    _, err = wl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}
