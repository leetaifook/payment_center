package main

import (
    "fmt"
    "payment_center/core/transaction"
    "time"
)

func coreAccountReceivables(args map[string]interface{}, result *Result) {
    payeeid := args["payeeid"].(float64)
    payerid := args["payerid"].(float64)
    amount := args["amount"].(float64)
    r, err := transaction.NewReceivables(int64(payeeid), int64(payerid), int64(amount))
    if err != nil {
        result.Res = 0
        result.Error = fmt.Sprintf("%v", err)
        result.Time = time.Now().Unix()
    } else {
        result.Res = 1
        result.Data = map[string]int64{"id": r.Id}
        result.Time = time.Now().Unix()
    }
}
