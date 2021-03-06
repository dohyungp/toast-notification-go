package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

const DOMAIN = "https://api-sms.cloud.toast.com"
const DEFAULT_API_VERSION = "v3.0"
const DEFAULT_TIMEOUT = 10 * time.Second
const DEFAULT_MAX_CONNCURRENCY = 10

type ToastClient struct {
	AppKey        string
	ApiSecret     string
	ApiVersion    string
	MaxConcurreny int
	Validator     *validator.Validate
	Client        *http.Client
}

type ExtraConfig struct {
	TimoutSecond  uint
	MaxConcurreny int
	ApiVersion    string
}

// ToastClient를 생성한다.
func NewToastClient(AppKey string, ApiSecret string, extras ...*ExtraConfig) *ToastClient {

	timeoutSecond := DEFAULT_TIMEOUT
	maxConcurrency := DEFAULT_MAX_CONNCURRENCY
	apiVersion := DEFAULT_API_VERSION

	if len(extras) != 0 {
		c := extras[0]
		if c.TimoutSecond != 0 {
			timeoutSecond = time.Duration(c.TimoutSecond) * time.Second
		}

		if c.MaxConcurreny != 0 {
			maxConcurrency = c.MaxConcurreny
		}

		if c.ApiVersion != "" {
			apiVersion = c.ApiVersion
		}
	}

	return &ToastClient{
		AppKey:     AppKey,
		ApiSecret:  ApiSecret,
		ApiVersion: apiVersion,
		Validator:  validator.New(),
		Client: &http.Client{
			Timeout: timeoutSecond,
		},
		MaxConcurreny: maxConcurrency,
	}
}

// Toast 요청 포맷에 맞게 Request를 준비한다.
func (t ToastClient) prepareRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json;charset=UTF-8"},
		"X-Secret-Key": []string{t.ApiSecret},
	}

	return req, nil
}

// 코드를 Validation한다.
func (t ToastClient) Validate(schema interface{}) {
	err := t.Validator.Struct(schema)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
}

// Toast로 SMS 메시지를 보낸다
func (t ToastClient) SendMessage(message TextMessage) {
	t.Validate(message)
	url := fmt.Sprintf("%s/sms/%s/appKeys/%s/sender/%s", DOMAIN, t.ApiVersion, t.AppKey, message.Type)
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	var sendResult SendResponse
	err = json.Unmarshal(body, &sendResult)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	fmt.Printf("Results: %v\n", sendResult)
}
