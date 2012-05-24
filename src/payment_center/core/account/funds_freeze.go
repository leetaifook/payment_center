package account

import (
    "payment_center/mysql"
    "strconv"
    "time"
)

type FundsFreeze struct {
    Id         int64
    AccountId  int64
    Type       byte
    Status     byte
    Amount     int64
    Enabled    byte
    Reason     string
    CreateTime int64
    UpdateTime int64
}

func (ff *FundsFreeze) Create() error {
    if ff.Id > 0 {
        return setError("已经存在此资金冻结信息")
    }

    a := &Account{Id: ff.AccountId}
    err := a.GetNoFreeze()
    if err != nil {
        return err
    }

    typeText := ""
    if ff.Type == 0 {
        typeText = "解冻"
        if a.FreezeAmount < ff.Amount {
            return setError("已冻结资金小于需" + typeText + "资金")
        }
    } else {
        typeText = "冻结"
        if a.UseAmount < ff.Amount {
            return setError("可使用资金小于需" + typeText + "资金")
        }
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "funds_freeze SET " +
        "account_id=?, " +
        "type=?, " +
        "status=?, " +
        "amount=?, " +
        "enabled=?, " +
        "reason=?, " +
        "create_time=?, " +
        "update_time=?")
    if err != nil {
        return err
    }

    now := time.Now().Unix()
    ff.CreateTime = now
    ff.UpdateTime = now
    res, err := ins.Exec(
        ff.AccountId,
        ff.Type,
        ff.Status,
        ff.Amount,
        ff.Enabled,
        ff.Reason,
        ff.CreateTime,
        ff.UpdateTime,
    )
    if err != nil {
        return err
    }

    ffid, err := res.LastInsertId()
    ff.Id = ffid

    ffl := &FundsFreezeLog{
        FundsFreezeId: ff.Id,
        AccountId:     ff.AccountId,
        Amount:        ff.Amount,
        TypeCode:      ff.Type,
        Type:          typeText,
        StatusCode:    1,
        Status:        "成功",
        Memo:          "新建：对账户编号：" + strconv.FormatInt(ff.AccountId, 10) + typeText + "(" + strconv.FormatInt(int64(ff.Type), 10) + ")资金" + strconv.FormatInt(ff.Amount, 10) + "，操作编号：" + strconv.FormatInt(ff.Id, 10) + ",理由为：" + ff.Reason,
    }

    _, err = ffl.Create()
    if err != nil {
        return err
    } else {
        ff.Enabled = 1
        _, err := ff.Update()
        if err != nil {
            return err
        }
    }

    al := &AccountLog{
        AccountId:     ff.AccountId,
        OptTypeCode:   7,
        OptType:       "资金冻结",
        OptId:         ff.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "对账户编号：" + strconv.FormatInt(ff.AccountId, 10) + typeText,
    }
    _, err = al.Create()
    if err != nil {
        return err
    }

    if ff.Type == 0 {
        a.UseAmount += ff.Amount
        a.WithdrawalsAmount += ff.Amount
        a.FreezeAmount -= ff.Amount
    } else {
        a.UseAmount -= ff.Amount
        a.WithdrawalsAmount -= ff.Amount
        a.FreezeAmount += ff.Amount
    }

    _, err = a.Update()
    if err != nil {
        return err
    }

    return nil
}

func (ff *FundsFreeze) Update() (int64, error) {
    if ff.Id <= 0 {
        return 0, setError("无此资金冻结信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "funds_freeze SET " +
        "account_id=?, " +
        "type=?, " +
        "status=?, " +
        "amount=?, " +
        "enabled=?, " +
        "reason=?, " +
        "update_time=? " +
        "WHERE id=?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    ff.UpdateTime = now
    res, err := upd.Exec(
        ff.AccountId,
        ff.Type,
        ff.Status,
        ff.Amount,
        ff.Enabled,
        ff.Reason,
        ff.UpdateTime,
        ff.Id,
    )
    if err != nil {
        return 0, err
    }

    typeText := ""
    if ff.Type == 0 {
        typeText = "解冻"
    } else {
        typeText = "冻结"
    }

    rowsAN, err := res.RowsAffected()
    ffl := &FundsFreezeLog{
        FundsFreezeId: ff.Id,
        AccountId:     ff.AccountId,
        Amount:        ff.Amount,
        TypeCode:      ff.Type,
        Type:          typeText,
        StatusCode:    1,
        Status:        "成功",
        Memo:          "更新：对账户编号：" + strconv.FormatInt(ff.AccountId, 10) + typeText + "(" + strconv.FormatInt(int64(ff.Type), 10) + ")资金" + strconv.FormatInt(ff.Amount, 10) + "，操作编号：" + strconv.FormatInt(ff.Id, 10) + ",理由为：" + ff.Reason,
    }
    _, err = ffl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}
