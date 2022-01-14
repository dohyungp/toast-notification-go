package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dohyungp/toast-notification-go/schema"
	"github.com/go-playground/validator/v10"
)

const DOMAIN = "https://api-sms.cloud.toast.com"
const API_VERSION = "v3.0"

type ToastClient struct {
	AppKey    string
	ApiSecret string
	Validator *validator.Validate
}

// ToastClient를 생성한다.
func NewToastClient(AppKey string, ApiSecret string) *ToastClient {
	return &ToastClient{
		AppKey:    AppKey,
		ApiSecret: ApiSecret,
		Validator: validator.New(),
	}
}

func (t ToastClient) Validate(schema interface{}) {
	err := t.Validator.Struct(schema)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
}

// Toast로 SMS 메시지를 보낸다
func (t ToastClient) SendMessage(message schema.TextMessage) {
	url := fmt.Sprintf("%s/sms/%s/appKeys/%s/sender/sms", DOMAIN, API_VERSION, t.AppKey)
	t.Validate(message)
	// FIXME: _를 에러 핸들링처리하도록 로직 변경이 필요하다.
	msgJson, _ := json.Marshal(message)
	msgBuffer := bytes.NewBuffer(msgJson)
	// FIXME: _를 에러 핸들링처리하도록 로직 변경이 필요하다.
	resp, _ := http.Post(url, "application/json", msgBuffer)
	// FIXME: Response를 처리하는 로직이 필요하다.
	resp.Body.Close()
}
