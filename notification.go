package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dohyungp/toast-notification-go/schema"
	"github.com/go-playground/validator/v10"
)

const DOMAIN = "https://api-sms.cloud.toast.com"
const API_VERSION = "v3.0"

type ToastClient struct {
	AppKey    string
	ApiSecret string
	Validator *validator.Validate
	Client    *http.Client
}

type Config struct {
	TimoutSecond uint
}

// ToastClient를 생성한다.
func NewToastClient(AppKey string, ApiSecret string, config ...*Config) *ToastClient {

	timeoutSecond := 10 * time.Second

	for _, c := range config {
		if c.TimoutSecond != 0 {
			timeoutSecond = time.Duration(c.TimoutSecond) * time.Second
		}
		// NOTE: 최초 건에 대해서만 적용하도록 break 처리한다.
		break
	}

	return &ToastClient{
		AppKey:    AppKey,
		ApiSecret: ApiSecret,
		Validator: validator.New(),
		Client: &http.Client{
			Timeout: timeoutSecond,
		},
	}
}

// Toast 요청 포맷에 맞게 Request를 준비한다.
func (t ToastClient) prepareRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"X-Secret-Key": []string{t.ApiSecret},
	}

	return req, nil
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
	req, err := t.prepareRequest("POST", url, msgBuffer)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(body)
}
