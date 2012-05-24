package account

import (
    "payment_center/mysql"
    "strconv"
    "time"
)

type Recharge struct {
    Id         int64
    AccountId  int64
    Amount     int64
    Status     byte
    Enabled    byte
    CreateTime int64
    UpdateTime int64
}

func (r *Recharge) Create() error {
    if r.Id > 0 {
        return setError("已经存在此充值信息")
    }

    a := &Account{Id: r.AccountId}
    err := a.GetNoFreeze()
    if err != nil {
        return err
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "recharge SET " +
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
    r.CreateTime = now
    r.UpdateTime = now
    res, err := ins.Exec(
        r.AccountId,
        r.Amount,
        r.Status,
        r.Enabled,
        r.CreateTime,
        r.UpdateTime,
    )
    if err != nil {
        return err
    }

    rid, err := res.LastInsertId()
    r.Id = rid
    rl := &RechargeLog{
        RechargeId: r.Id,
        AccountId:  r.AccountId,
        Amount:     r.Amount,
        StatusCode: 1,
        Status:     "成功",
        Memo:       "新建：对账户编号：" + strconv.FormatInt(r.AccountId, 10) + "充值，充值编号：" + strconv.FormatInt(r.Id, 10) + ",充值金额为：" + strconv.FormatInt(r.Amount, 10),
    }

    _, err = rl.Create()
    if err != nil {
        return err
    } else {
        r.Enabled = 1
        _, err := r.Update()
        if err != nil {
            return err
        }
    }

    al := &AccountLog{
        AccountId:     r.AccountId,
        OptTypeCode:   3,
        OptType:       "充值",
        OptId:         r.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "充值账户编号：" + strconv.FormatInt(r.AccountId, 10),
    }
    _, err = al.Create()
    if err != nil {
        return err
    }

    a.TotalAmount += r.Amount
    a.UseAmount += r.Amount
    _, err = a.Update()
    if err != nil {
        return err
    }

    return nil
}

func (r *Recharge) Update() (int64, error) {
    if r.Id <= 0 {
        return 0, setError("无此充值信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "recharge SET " +
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
    r.UpdateTime = now
    res, err := upd.Exec(
        r.AccountId,
        r.Amount,
        r.Status,
        r.Enabled,
        r.UpdateTime,
        r.Id,
    )
    if err != nil {
        return 0, err
    }

    rowsAN, err := res.RowsAffected()
    rl := &RechargeLog{
        RechargeId: r.Id,
        AccountId:  r.AccountId,
        Amount:     r.Amount,
        StatusCode: 1,
        Status:     "成功",
        Memo:       "更新：对账户编号：" + strconv.FormatInt(r.AccountId, 10) + "充值，充值编号：" + strconv.FormatInt(r.Id, 10) + ",充值金额为：" + strconv.FormatInt(r.Amount, 10),
    }
    _, err = rl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}
