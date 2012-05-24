package account

import (
    "encoding/json"
    "fmt"
    "time"
)

type AccountError struct {
    What string `json:"what"`
}

func (e *AccountError) Error() string {
    m := map[string]interface{}{
        "when": time.Now().Unix(),
        "what": e.What,
    }
    jsonByte, _ := json.Marshal(m)
    return fmt.Sprintf("%s", jsonByte)
}
