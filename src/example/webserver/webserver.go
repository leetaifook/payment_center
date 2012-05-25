package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
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
                coreAccountCreate(args, result)
            case "core.account.Info":
                coreAccountInfo(args, result)
            case "core.account.AccountFreeze":
                coreAccountAccountFreeze(args, result)
            case "core.account.FundsFreeze":
                coreAccountFundsFreeze(args, result)
            case "core.account.Recharge":
                coreAccountRecharge(args, result)
            case "core.account.Withdrawals":
                coreAccountWithdrawals(args, result)
            case "core.account.Transfer":
                coreAccountTransfer(args, result)
            case "transaction.Payment":
                coreAccountPayment(args, result)
            case "transaction.Receivables":
                coreAccountReceivables(args, result)
            default:
                result.Res = 0
                result.Error = "无此命令:" + fmt.Sprintf("%v", cmd["func"])
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
