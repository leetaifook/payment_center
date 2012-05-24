package account

import (
    "payment_center/mysql"
    "strconv"
    "time"
)

type Transfer struct {
    Id             int64
    PayerAccountId int64
    PayeeAccountId int64
    Amount         int64
    Status         byte
    Enabled        byte
    CreateTime     int64
    UpdateTime     int64
}

func (t *Transfer) Create() error {
    if t.Id > 0 {
        return setError("已经存在此转账信息")
    }

    payer := &Account{Id: t.PayerAccountId}
    err := payer.GetNoFreeze()
    if err != nil {
        return err
    }

    payee := &Account{Id: t.PayeeAccountId}
    err = payee.GetNoFreeze()
    if err != nil {
        return err
    }

    if payer.Currency != payee.Currency {
        return setError("出账方与入账方账户币种不一致")
    }

    if payer.UseAmount < t.Amount {
        return setError("可使用金额少于转账金额")
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "transfer SET " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "amount=?, " +
        "status=?, " +
        "enabled=?, " +
        "create_time=?, " +
        "update_time=?")
    if err != nil {
        return err
    }

    now := time.Now().Unix()
    t.CreateTime = now
    t.UpdateTime = now
    res, err := ins.Exec(
        t.PayerAccountId,
        t.PayeeAccountId,
        t.Amount,
        t.Status,
        t.Enabled,
        t.CreateTime,
        t.UpdateTime,
    )
    if err != nil {
        return err
    }

    tid, err := res.LastInsertId()
    t.Id = tid
    tl := &TransferLog{
        TransferId:     t.Id,
        PayerAccountId: t.PayerAccountId,
        PayeeAccountId: t.PayeeAccountId,
        Amount:         t.Amount,
        StatusCode:     1,
        Status:         "成功",
        Memo:           "新建：出账方编号：" + strconv.FormatInt(t.PayerAccountId, 10) + "向入账方编号：" + strconv.FormatInt(t.PayeeAccountId, 10) + "转账，转账编号：" + strconv.FormatInt(t.Id, 10) + ",转账金额为：" + strconv.FormatInt(t.Amount, 10),
    }

    _, err = tl.Create()
    if err != nil {
        return err
    } else {
        t.Enabled = 1
        _, err := t.Update()
        if err != nil {
            return err
        }
    }

    payerLog := &AccountLog{
        AccountId:     t.PayerAccountId,
        OptTypeCode:   5,
        OptType:       "转账",
        OptId:         t.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "出账方账户编号：" + strconv.FormatInt(t.PayerAccountId, 10) + ",入账方账户编号：" + strconv.FormatInt(t.PayeeAccountId, 10),
    }
    _, err = payerLog.Create()
    if err != nil {
        return err
    }

    payer.TotalAmount -= t.Amount
    payer.UseAmount -= t.Amount

    if payer.UseAmount-payer.WithdrawalsAmount <= t.Amount {
        payer.WithdrawalsAmount = 0
    } else {
        payer.WithdrawalsAmount -= t.Amount
    }

    _, err = payer.Update()
    if err != nil {
        return err
    }

    payeeLog := &AccountLog{
        AccountId:     t.PayeeAccountId,
        OptTypeCode:   5,
        OptType:       "转账",
        OptId:         t.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "入账方账户编号：" + strconv.FormatInt(t.PayeeAccountId, 10) + ",出账方账户编号：" + strconv.FormatInt(t.PayerAccountId, 10),
    }
    _, err = payeeLog.Create()
    if err != nil {
        return err
    }

    payee.TotalAmount += t.Amount
    payee.UseAmount += t.Amount
    payee.WithdrawalsAmount += t.Amount

    _, err = payee.Update()
    if err != nil {
        return err
    }

    return nil
}

func (t *Transfer) Update() (int64, error) {
    if t.Id <= 0 {
        return 0, setError("无此转账信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "transfer SET " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "amount=?, " +
        "status=?, " +
        "enabled=?, " +
        "update_time=? " +
        "WHERE id = ?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    t.UpdateTime = now
    res, err := upd.Exec(
        t.PayerAccountId,
        t.PayeeAccountId,
        t.Amount,
        t.Status,
        t.Enabled,
        t.UpdateTime,
        t.Id,
    )
    if err != nil {
        return 0, err
    }

    rowsAN, err := res.RowsAffected()
    tl := &TransferLog{
        TransferId:     t.Id,
        PayerAccountId: t.PayerAccountId,
        PayeeAccountId: t.PayeeAccountId,
        Amount:         t.Amount,
        StatusCode:     1,
        Status:         "成功",
        Memo:           "更新：出账方编号：" + strconv.FormatInt(t.PayerAccountId, 10) + "向入账方编号：" + strconv.FormatInt(t.PayeeAccountId, 10) + "转账，转账编号：" + strconv.FormatInt(t.Id, 10) + ",转账金额为：" + strconv.FormatInt(t.Amount, 10),
    }
    _, err = tl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}
