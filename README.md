# Go Toast Notification Client

발송 코드 예시

```go
package main

import (
    toast "github.com/dohyungp/toast-notification-go"
    "github.com/dohyungp/toast-notification-go/schema"
)

func main() {
	client := toast.NewToastClient("<APP_KEY>", "<API_SECRET>")
	recipient := schema.Recipient{RecipientNo: "01000000000"}
	message := schema.TextMessage{
		Body:          "Hello, Toast",
		SendNo:        "0700000000",
		RecipientList: []schema.Recipient{recipient},
	}
	client.SendMessage(message)
}
```