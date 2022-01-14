# Go Toast Notification Client

발송 코드 예시

```go
package main

import (
    "github.com/dohyungp/toast-notification-go/sms"
)

func main() {
	client := sms.NewToastClient("<APP_KEY>", "<API_SECRET>")
	recipient := sms.Recipient{RecipientNo: "01000000000"}
	message := sms.TextMessage{
		Body:          "Hello, Toast",
		SendNo:        "0700000000",
		RecipientList: []sms.Recipient{recipient},
	}
	client.SendMessage(message)
}
```