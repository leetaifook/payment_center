package main

import (
    "fmt"
    "payment_center/core/account"
    "time"
)

func coreAccountWithdrawals(args map[string]interface{}, result *Result) {
    aid := args["aid"].(float64)
    amount := args["amount"].(float64)
    w, err := account.NewWithdrawals(int64(aid), int64(amount))
    if err != nil {
        result.Res = 0
        result.Error = fmt.Sprintf("%v", err)
        result.Time = time.Now().Unix()
    } else {
        result.Res = 1
        result.Data = map[string]int64{"id": w.Id}
        result.Time = time.Now().Unix()
    }
}
