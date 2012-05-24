package transaction

import (
    "payment_center/mysql"
    "time"
)

type TransactionLog struct {
    Id               int64
    TransactionId    int64
    TriggerId        int64
    PayerAccountId   int64
    PayeeAccountId   int64
    MethodCode       byte
    Method           string
    PayMethodCode    byte
    PayMethod        string
    Amount           int64
    TypeCode         byte
    Type             string
    StatusCode       byte
    Status           string
    SettleMethodCode byte
    SettleMethod     string
    SettleStatusCode byte
    SettleStatus     string
    Memo             string
    StartTime        int64
    EndTime          int64
    CreateTime       int64
}

func (tl *TransactionLog) Create() (int64, error) {
    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "transaction_log SET " +
        "transaction_id=?, " +
        "trigger_id=?, " +
        "payer_account_id=?, " +
        "payee_account_id=?, " +
        "method_code=?, " +
        "method=?, " +
        "pay_method_code=?, " +
        "pay_method=?, " +
        "amount=?, " +
        "type_code=?, " +
        "type=?, " +
        "status_code=?, " +
        "status=?, " +
        "settle_method_code=?, " +
        "settle_method=?, " +
        "settle_status_code=?, " +
        "settle_status=?, " +
        "memo=?, " +
        "start_time=?, " +
        "end_time=?, " +
        "create_time=?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    tl.CreateTime = now
    res, err := ins.Exec(
        tl.TransactionId,
        tl.TriggerId,
        tl.PayerAccountId,
        tl.PayeeAccountId,
        tl.MethodCode,
        tl.Method,
        tl.PayMethodCode,
        tl.PayMethod,
        tl.Amount,
        tl.TypeCode,
        tl.Type,
        tl.StatusCode,
        tl.Status,
        tl.SettleMethodCode,
        tl.SettleMethod,
        tl.SettleStatusCode,
        tl.SettleStatus,
        tl.Memo,
        tl.StartTime,
        tl.EndTime,
        tl.CreateTime,
    )
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}
