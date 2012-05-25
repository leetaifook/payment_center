package main

import (
    "fmt"
    "payment_center/core/account"
    "reflect"
    "time"
)

func coreAccountInfo(args map[string]interface{}, result *Result) {
    aid := args["aid"].(float64)
    a, err := account.AccountInfo(int64(aid))
    m := make(map[string]interface{})
    v := reflect.ValueOf(a)
    t := v.Type()
    for i := 0; i < t.NumField(); i++ {
        if t.Field(i).Name != "Password" {
            f := v.Field(i)
            m[t.Field(i).Name] = f.Interface()
        }
    }

    if err != nil {
        result.Res = 0
        result.Error = fmt.Sprintf("%v", err)
        result.Time = time.Now().Unix()
    } else {
        result.Res = 1
        result.Data = m
        result.Time = time.Now().Unix()
    }
}
