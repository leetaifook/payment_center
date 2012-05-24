package main

import (
    "fmt"
    "payment_center/core/transaction"
    "time"
)

func coreAccountPayment(args map[string]interface{}, result *Result) {
    payerid := args["payerid"].(float64)
    payeeid := args["payeeid"].(float64)
    amount := args["amount"].(float64)
    p, err := transaction.NewPayment(int64(payerid), int64(payeeid), int64(amount))
    if err != nil {
        result.Res = 0
        result.Error = fmt.Sprintf("%v", err)
        result.Time = time.Now().Unix()
    } else {
        result.Res = 1
        result.Data = map[string]int64{"id": p.Id}
        result.Time = time.Now().Unix()
    }
}
