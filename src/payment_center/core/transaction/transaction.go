package transaction

import (
    "database/sql"
    "payment_center/mysql"
    //"strconv"
    "time"
)

type Transaction struct {
    Id             int64
    TriggerId      int64
    PayerAccountId int64
    PayeeAccountId int64
    Method         byte
    PayMethod      byte
    Amount         int64
    Type           byte
    Status         byte
    SettleMethod   byte
    SettleStatus   byte
    SettleTime     int64
    Enabled        byte
    StartTime      int64
    EndTime        int64
    CreateTime     int64
    UpdateTime     int64
}

func (t *Transaction) Create() error {
    if t.Id > 0 {
        return setError("已经存在此交易信息")
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "transaction SET " +
        "trigger_id=?, " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "method=?, " +
        "pay_method=?, " +
        "amount=?, " +
        "type=?, " +
        "status=?, " +
        "settle_method=?, " +
        "settle_status=?, " +
        "settle_time=?, " +
        "enabled=?, " +
        "start_time=?, " +
        "end_time=?, " +
        "create_time=?, " +
        "update_time=?")
    if err != nil {
        return err
    }

    now := time.Now().Unix()
    t.CreateTime = now
    t.UpdateTime = now
    res, err := ins.Exec(
        t.TriggerId,
        t.PayerAccountId,
        t.PayeeAccountId,
        t.Method,
        t.PayMethod,
        t.Amount,
        t.Type,
        t.Status,
        t.SettleMethod,
        t.SettleStatus,
        t.SettleTime,
        t.Enabled,
        t.StartTime,
        t.EndTime,
        t.CreateTime,
        t.UpdateTime,
    )
    if err != nil {
        return err
    }

    tid, err := res.LastInsertId()
    t.Id = tid
    tl := &TransactionLog{
        TransactionId:    t.Id,
        TriggerId:        t.TriggerId,
        PayerAccountId:   t.PayerAccountId,
        PayeeAccountId:   t.PayeeAccountId,
        MethodCode:       t.Method,
        Method:           "",
        PayMethodCode:    t.PayMethod,
        PayMethod:        "",
        Amount:           t.Amount,
        TypeCode:         t.Type,
        Type:             "",
        StatusCode:       t.Status,
        Status:           "",
        SettleMethodCode: t.SettleMethod,
        SettleMethod:     "",
        SettleStatusCode: t.SettleStatus,
        SettleStatus:     "",
        Memo:             "新建：",
        StartTime:        t.StartTime,
        EndTime:          t.EndTime,
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

    return nil
}

func (t *Transaction) Update() (int64, error) {
    if t.Id <= 0 {
        return 0, setError("无此交易信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "transaction SET " +
        "trigger_id=?, " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "method=?, " +
        "pay_method=?, " +
        "amount=?, " +
        "type=?, " +
        "status=?, " +
        "settle_method=?, " +
        "settle_status=?, " +
        "settle_time=?, " +
        "enabled=?, " +
        "start_time=?, " +
        "end_time=?, " +
        "update_time=? " +
        "WHERE id = ?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    t.UpdateTime = now
    res, err := upd.Exec(
        t.TriggerId,
        t.PayerAccountId,
        t.PayeeAccountId,
        t.Method,
        t.PayMethod,
        t.Amount,
        t.Type,
        t.Status,
        t.SettleMethod,
        t.SettleStatus,
        t.SettleTime,
        t.Enabled,
        t.StartTime,
        t.EndTime,
        t.UpdateTime,
        t.Id,
    )
    if err != nil {
        return 0, err
    }

    rowsAN, err := res.RowsAffected()
    tl := &TransactionLog{
        TransactionId:    t.Id,
        TriggerId:        t.TriggerId,
        PayerAccountId:   t.PayerAccountId,
        PayeeAccountId:   t.PayeeAccountId,
        MethodCode:       t.Method,
        Method:           "",
        PayMethodCode:    t.PayMethod,
        PayMethod:        "",
        Amount:           t.Amount,
        TypeCode:         t.Type,
        Type:             "",
        StatusCode:       t.Status,
        Status:           "",
        SettleMethodCode: t.SettleMethod,
        SettleMethod:     "",
        SettleStatusCode: t.SettleStatus,
        SettleStatus:     "",
        Memo:             "更新：",
        StartTime:        t.StartTime,
        EndTime:          t.EndTime,
    }
    _, err = tl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}

func (t *Transaction) Get() error {
    var row *sql.Row

    if t.Id > 0 {
        row = doDb.QueryRow("SELECT * FROM "+mysql.PreTable+"transaction WHERE id =? LIMIT 1", t.Id)
    } else {
        if t.TriggerId <= 0 {
            return setError("无触发交易行为编号")
        }

        if t.Method <= 0 {
            return setError("无交易方式")
        }

        row = doDb.QueryRow("SELECT * FROM "+mysql.PreTable+"transaction WHERE trigger_id =? AND method =? LIMIT 1", t.TriggerId, t.Method)
    }
    err := row.Scan(
        &t.Id,
        &t.TriggerId,
        &t.PayerAccountId,
        &t.PayeeAccountId,
        &t.Method,
        &t.PayMethod,
        &t.Amount,
        &t.Type,
        &t.Status,
        &t.SettleMethod,
        &t.SettleStatus,
        &t.SettleTime,
        &t.Enabled,
        &t.StartTime,
        &t.EndTime,
        &t.CreateTime,
        &t.UpdateTime,
    )

    if err != nil {
        return setError("无此交易")
    }

    return nil
}
