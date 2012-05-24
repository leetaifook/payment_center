package main

import (
    "fmt"
    "payment_center/core/account"
    "time"
)

func coreAccountAccountFreeze(args map[string]interface{}, result *Result) {
    aid := args["aid"].(float64)
    freeze := args["freeze"].(float64)
    reason := args["reason"].(string)
    af, err := account.NewAccountFreeze(int64(aid), byte(freeze), reason)
    if err != nil {
        result.Res = 0
        result.Error = fmt.Sprintf("%v", err)
        result.Time = time.Now().Unix()
    } else {
        result.Res = 1
        result.Data = map[string]int64{"id": af.Id}
        result.Time = time.Now().Unix()
    }
}
