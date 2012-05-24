package account

import (
    "payment_center/mysql"
    "strconv"
    "time"
)

type AccountFreeze struct {
    Id         int64
    AccountId  int64
    Type       byte
    Status     byte
    Enabled    byte
    Reason     string
    CreateTime int64
    UpdateTime int64
}

func (af *AccountFreeze) Create() error {
    if af.Id > 0 {
        return setError("已经存在此账户冻结信息")
    }

    a := &Account{Id: af.AccountId}
    err := a.Get()
    if err != nil {
        return err
    }

    if a.Freeze == af.Type {
        return setError("此账户已经处于需要设置的此冻结状态")
    }

    ins, err := doDb.Prepare("INSERT INTO " + mysql.PreTable + "account_freeze SET " +
        "account_id=?, " +
        "type=?, " +
        "status=?, " +
        "enabled=?, " +
        "reason=?, " +
        "create_time=?, " +
        "update_time=?")
    if err != nil {
        return err
    }

    now := time.Now().Unix()
    af.CreateTime = now
    af.UpdateTime = now
    res, err := ins.Exec(
        af.AccountId,
        af.Type,
        af.Status,
        af.Enabled,
        af.Reason,
        af.CreateTime,
        af.UpdateTime,
    )
    if err != nil {
        return err
    }

    afid, err := res.LastInsertId()
    af.Id = afid

    typeText := ""
    if af.Type == 0 {
        typeText = "解冻"
    } else {
        typeText = "冻结"
    }

    afl := &AccountFreezeLog{
        AccountFreezeId: af.Id,
        AccountId:       af.AccountId,
        TypeCode:        af.Type,
        Type:            typeText,
        StatusCode:      1,
        Status:          "成功",
        Memo:            "新建：对账户编号：" + strconv.FormatInt(af.AccountId, 10) + typeText + "(" + strconv.FormatInt(int64(af.Type), 10) + ")，操作编号：" + strconv.FormatInt(afid, 10) + ",理由为：" + af.Reason,
    }

    _, err = afl.Create()
    if err != nil {
        return err
    } else {
        af.Enabled = 1
        _, err := af.Update()
        if err != nil {
            return err
        }
    }

    al := &AccountLog{
        AccountId:     af.AccountId,
        OptTypeCode:   6,
        OptType:       "账户冻结",
        OptId:         af.Id,
        OptStatusCode: 1,
        OptStatus:     "成功",
        Memo:          "对账户编号：" + strconv.FormatInt(af.AccountId, 10) + typeText,
    }
    _, err = al.Create()
    if err != nil {
        return err
    }

    a.Freeze = af.Type
    _, err = a.Update()
    if err != nil {
        return err
    }

    return nil
}

func (af *AccountFreeze) Update() (int64, error) {
    if af.Id <= 0 {
        return 0, setError("无此账户冻结信息")
    }

    upd, err := doDb.Prepare("UPDATE " + mysql.PreTable + "account_freeze SET " +
        "account_id=?, " +
        "type=?, " +
        "status=?, " +
        "enabled=?, " +
        "reason=?, " +
        "update_time=? " +
        "WHERE id=?")
    if err != nil {
        return 0, err
    }

    now := time.Now().Unix()
    af.UpdateTime = now
    res, err := upd.Exec(
        af.AccountId,
        af.Type,
        af.Status,
        af.Enabled,
        af.Reason,
        af.UpdateTime,
        af.Id,
    )
    if err != nil {
        return 0, err
    }

    typeText := ""
    if af.Type == 0 {
        typeText = "解冻"
    } else {
        typeText = "冻结"
    }

    rowsAN, err := res.RowsAffected()
    afl := &AccountFreezeLog{
        AccountFreezeId: af.Id,
        AccountId:       af.AccountId,
        TypeCode:        af.Type,
        Type:            typeText,
        StatusCode:      1,
        Status:          "成功",
        Memo:            "更新：对账户编号：" + strconv.FormatInt(af.AccountId, 10) + typeText + "(" + strconv.FormatInt(int64(af.Type), 10) + ")，操作编号：" + strconv.FormatInt(af.Id, 10) + ",理由为：" + af.Reason,
    }
    _, err = afl.Create()
    if err != nil {
        return 0, err
    }

    return rowsAN, nil
}
