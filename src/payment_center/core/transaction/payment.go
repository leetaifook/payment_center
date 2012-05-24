package transaction

import (
    "payment_center/core/account"
    "payment_center/mysql"
    //"strconv"
    "time"
)

type Payment struct {
    Id             int64
    TransferId     int64
    PayerAccountId int64
    PayeeAccountId int64
    Amount         int64
    Type           byte
    Status         byte
    Currency       byte
    Enabled        byte
    CreateTime     int64
    UpdateTime     int64
}

func (p *Payment) Create() error {
    if p.Id > 0 {
        return setError("已经存在此付款信息")
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "payment SET " +
        "transfer_id=?, " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "amount=?, " +
        "type=?, " +
        "status=?, " +
        "currency=?, " +
        "enabled=?, " +
        "create_time=?, " +
        "update_time=?")
    if err != nil {
        return err
    }

    now := time.Now().Unix()
    p.CreateTime = now
    p.UpdateTime = now
    res, err := ins.Exec(
        p.TransferId,
        p.PayerAccountId,
        p.PayeeAccountId,
        p.Amount,
        p.Type,
        p.Status,
        p.Currency,
        p.Enabled,
        p.CreateTime,
        p.UpdateTime,
    )
    if err != nil {
        return err
    }

    pid, err := res.LastInsertId()
    p.Id = pid
    pl := &PaymentLog{
        PaymentId:    p.Id,
        Amount:       p.Amount,
        TypeCode:     p.Type,
        Type:         "",
        StatusCode:   p.Status,
        Status:       "成功",
        CurrencyCode: p.Currency,
        Currency:     "人民币",
        Memo:         "新建：",
    }

    _, err = pl.Create()
    if err != nil {
        return err
    } else {
        p.Enabled = 1
        _, err := p.Update()
        if err != nil {
            return err
        }
    }

    t := &Transaction{
        TriggerId:      p.Id,
        PayerAccountId: p.PayerAccountId,
        PayeeAccountId: p.PayeeAccountId,
        Method:         1,
        PayMethod:      1,
        Amount:         p.Amount,
        Type:           p.Type,
        Status:         p.Status,
        SettleMethod:   1,
        SettleStatus:   1,
        StartTime:      now,
        EndTime:        (now + 7*24*60*60),
    }

    err = t.Create()
    if err != nil {
        return err
    }

    transfer, err := account.NewTransfer(p.PayerAccountId, p.PayeeAccountId, p.Amount)
    if err != nil {
        return err
    }

    p.TransferId = transfer.Id
    _, err = p.Update()
    if err != nil {
        return err
    }

    return nil
}

func (p *Payment) Update() (int64, error) {
    if p.Id <= 0 {
        return 0, setError("无此付款信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "payment SET " +
        "transfer_id=?, " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "amount=?, " +
        "type=?, " +
        "status=?, " +
        "currency=?, " +
        "enabled=?, " +
        "update_time=? " +
        "WHERE id = ?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    p.UpdateTime = now
    res, err := upd.Exec(
        p.TransferId,
        p.PayerAccountId,
        p.PayeeAccountId,
        p.Amount,
        p.Type,
        p.Status,
        p.Currency,
        p.Enabled,
        p.UpdateTime,
        p.Id,
    )
    if err != nil {
        return 0, err
    }

    rowsAN, err := res.RowsAffected()
    pl := &PaymentLog{
        PaymentId:    p.Id,
        Amount:       p.Amount,
        TypeCode:     p.Type,
        Type:         "",
        StatusCode:   p.Status,
        Status:       "成功",
        CurrencyCode: p.Currency,
        Currency:     "人民币",
        Memo:         "更新：",
    }
    _, err = pl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}
