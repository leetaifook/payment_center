package main

import (
    "fmt"
    "payment_center/core/account"
    "time"
)

func coreAccountCreate(args map[string]interface{}, result *Result) {
    password := args["password"].(string)
    a, err := account.NewAccount(password, 1)
    if err != nil {
        result.Res = 0
        result.Error = fmt.Sprintf("%v", err)
        result.Time = time.Now().Unix()
    } else {
        result.Res = 1
        result.Data = map[string]int64{"id": a.Id}
        result.Time = time.Now().Unix()
    }
}
