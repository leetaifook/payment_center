package transaction

import (
    "encoding/json"
    "fmt"
    "time"
)

type TransactionError struct {
    What string `json:"what"`
}

func (e *TransactionError) Error() string {
    m := map[string]interface{}{
        "when": time.Now().Unix(),
        "what": e.What,
    }
    jsonByte, _ := json.Marshal(m)
    return fmt.Sprintf("%s", jsonByte)
}
