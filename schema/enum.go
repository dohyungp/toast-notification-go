package schema

// 메시지 상태 코드
type MessageStatus string

// 수신 결과 코드
type ReceiveResult string

// 수신 결과 코드 상세
type ReceiveResultDetail string

const (
	MESSAGE_FAIL        MessageStatus = "0"
	MESSAGE_REQUEST     MessageStatus = "1"
	MESSAGE_IN_PROGRESS MessageStatus = "2"
	MESSAGE_SUCCESS     MessageStatus = "3"
	MESSAGE_CANCELED    MessageStatus = "4"
	MESSAGE_DUPLICATED  MessageStatus = "5"
)

const (
	RECEIVE_SUCCESS ReceiveResult = "MTR1"
	RECEIVE_FAIL    ReceiveResult = "MTR2"
)

const (
	RECEIVE_INVALID_REQUEST ReceiveResultDetail = "MTR2_1"
	RECEIVE_CARRIER_PROBLEM ReceiveResultDetail = "MTR2_2"
	RECEIVE_DEVICE_PROBLEM  ReceiveResultDetail = "MTR2_3"
)

type RequestStatus string

const (
	REQUEST_IN_PROGRESS RequestStatus = "1"
	REQUEST_COMPLETED   RequestStatus = "2"
	REQUEST_FAILED      RequestStatus = "3"
)
