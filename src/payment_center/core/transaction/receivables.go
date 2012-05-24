package transaction

import (
    "payment_center/core/account"
    "payment_center/mysql"
    //"strconv"
    "time"
)

type Receivables struct {
    Id             int64
    TransferId     int64
    PayeeAccountId int64
    PayerAccountId int64
    Amount         int64
    Type           byte
    Status         byte
    Currency       byte
    Enabled        byte
    CreateTime     int64
    UpdateTime     int64
}

func (r *Receivables) Create() error {
    if r.Id > 0 {
        return setError("已经存在此收款信息")
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "receivables SET " +
        "transfer_id=?, " +
        "payee_account_id=?, " +
        "payer_account_id=?, " +
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
    r.CreateTime = now
    r.UpdateTime = now
    res, err := ins.Exec(
        r.TransferId,
        r.PayeeAccountId,
        r.PayerAccountId,
        r.Amount,
        r.Type,
        r.Status,
        r.Currency,
        r.Enabled,
        r.CreateTime,
        r.UpdateTime,
    )
    if err != nil {
        return err
    }

    rid, err := res.LastInsertId()
    r.Id = rid
    rl := &ReceivablesLog{
        ReceivablesId: r.Id,
        Amount:        r.Amount,
        TypeCode:      r.Type,
        Type:          "",
        StatusCode:    r.Status,
        Status:        "成功",
        CurrencyCode:  r.Currency,
        Currency:      "人民币",
        Memo:          "新建：",
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

    t := &Transaction{
        TriggerId:      r.Id,
        PayerAccountId: r.PayerAccountId,
        PayeeAccountId: r.PayeeAccountId,
        Method:         2,
        PayMethod:      1,
        Amount:         r.Amount,
        Type:           r.Type,
        Status:         r.Status,
        SettleMethod:   1,
        SettleStatus:   1,
        StartTime:      now,
        EndTime:        (now + 7*24*60*60),
    }

    err = t.Create()

    if err != nil {
        return err
    }

    transfer, err := account.NewTransfer(r.PayerAccountId, r.PayeeAccountId, r.Amount)
    if err != nil {
        return err
    }

    r.TransferId = transfer.Id
    _, err = r.Update()
    if err != nil {
        return err
    }

    return nil
}

func (r *Receivables) Update() (int64, error) {
    if r.Id <= 0 {
        return 0, setError("无此收款信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "receivables SET " +
        "transfer_id=?, " +
        "payee_account_id=?, " +
        "payer_account_id=?, " +
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
    r.UpdateTime = now
    res, err := upd.Exec(
        r.TransferId,
        r.PayeeAccountId,
        r.PayerAccountId,
        r.Amount,
        r.Type,
        r.Status,
        r.Currency,
        r.Enabled,
        r.UpdateTime,
        r.Id,
    )
    if err != nil {
        return 0, err
    }

    rowsAN, err := res.RowsAffected()
    rl := &ReceivablesLog{
        ReceivablesId: r.Id,
        Amount:        r.Amount,
        TypeCode:      r.Type,
        Type:          "",
        StatusCode:    r.Status,
        Status:        "成功",
        CurrencyCode:  r.Currency,
        Currency:      "人民币",
        Memo:          "更新：",
    }
    _, err = rl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}
