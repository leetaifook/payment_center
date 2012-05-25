package main

import (
    "fmt"
    "payment_center/core/account"
    "time"
)

func coreAccountFundsFreeze(args map[string]interface{}, result *Result) {
    aid := args["aid"].(float64)
    amount := args["amount"].(float64)
    freeze := args["freeze"].(float64)
    reason := args["reason"].(string)
    ff, err := account.NewFundsFreeze(int64(aid), int64(amount), byte(freeze), reason)
    if err != nil {
        result.Res = 0
        result.Error = fmt.Sprintf("%v", err)
        result.Time = time.Now().Unix()
    } else {
        result.Res = 1
        result.Data = map[string]int64{"id": ff.Id}
        result.Time = time.Now().Unix()
    }
}
