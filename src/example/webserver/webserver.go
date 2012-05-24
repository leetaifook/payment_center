package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "payment_center/core/account"
    "payment_center/core/transaction"
    "runtime"
    "time"
)

type Result struct {
    Res   byte        `json:"result"`
    Data  interface{} `json:"data"`
    Error string      `json:"error"`
    Time  int64       `json:"time"`
}

func paymentCenter(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    result := new(Result)
    reqCmd := query.Get("cmd")
    if reqCmd == "" {
        result.Res = 0
        result.Error = "无此命令"
        result.Time = time.Now().Unix()
    } else {
        cmd := make(map[string]interface{})
        err := json.Unmarshal([]byte(reqCmd), &cmd)
        args := cmd["args"].(map[string]interface{})

        if err != nil {
            result.Res = 0
            result.Error = fmt.Sprintf("%v", err)
            result.Time = time.Now().Unix()
        } else {
            switch cmd["func"] {
            case "core.account.Create":
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

            case "core.account.AccountFreeze":
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
            case "core.account.Recharge":
                aid := args["aid"].(float64)
                amount := args["amount"].(float64)
                r, err := account.NewRecharge(int64(aid), int64(amount))
                if err != nil {
                    result.Res = 0
                    result.Error = fmt.Sprintf("%v", err)
                    result.Time = time.Now().Unix()
                } else {
                    result.Res = 1
                    result.Data = map[string]int64{"id": r.Id}
                    result.Time = time.Now().Unix()
                }
            case "core.account.Withdrawals":
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
            case "core.account.Transfer":
                payerid := args["payerid"].(float64)
                payeeid := args["payeeid"].(float64)
                amount := args["amount"].(float64)
                t, err := account.NewTransfer(int64(payerid), int64(payeeid), int64(amount))
                if err != nil {
                    result.Res = 0
                    result.Error = fmt.Sprintf("%v", err)
                    result.Time = time.Now().Unix()
                } else {
                    result.Res = 1
                    result.Data = map[string]int64{"id": t.Id}
                    result.Time = time.Now().Unix()
                }
            case "extension.account.Mapping":
                result.Res = 0
                result.Error = fmt.Sprintf("%v", err)
                result.Time = time.Now().Unix()
            case "transaction.Payment":
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
            case "transaction.Receivables":
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
            default:
                result.Res = 0
                result.Error = "无此命令"
                result.Time = time.Now().Unix()
            }
        }

    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    jsonByte, _ := json.Marshal(result)
    w.Write(jsonByte)
}

var addr = flag.String("addr", ":80", "Server port")

func main() {
    flag.Parse()
    runtime.GOMAXPROCS(runtime.NumCPU())
    http.HandleFunc("/api", paymentCenter)
    http.ListenAndServe(*addr, nil)
}
