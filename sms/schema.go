package sms

// Toast SMS API 수신인이다
// RecipientNo는 전화번호이며 필수이다
// CountryCode는 국가코드이며 없는 경우 자동으로 82가 된다
// InternationalRecipentNo는 국가코드를 포함한 전화번호이며 recipentNo가 있는 경우 무시된다
// TemplateParameter Toast에서 동적으로 문자 템플릿을 만들때 사용하며 Map 타입이다
// RecipientGroupingKey는 수신자 그룹키이다
type Recipient struct {
	RecipientNo              string                 `json:"recipientNo" validate:"max=50,required_without=InternationalRecipientNo"`
	CountryCode              string                 `json:"countryCode" validate:"max=8"`
	InternationalRecipientNo string                 `json:"internationalRecipientNo,omitempty" validate:"max=20,required_without=RecipientNo"`
	TemplateParameter        map[string]interface{} `json:"templateParameter,omitempty" validate:"omitempty,max=20"`
	RecipientGroupingKey     string                 `json:"recipientGroupingKey,omitempty" validate:"max=100"`
}

// Toast SMS 문자 메시지의 Request Body이다
// Body는 필수이나 TemplateId가 있는 경우에는 필수가 아니다.
// SendNo는 발신인 번호를 의미하며 필수이다.
// RequestDate는 yyyy-mm-dd HH:MM 포맷이며 발송 예약일시이다.
// SenderGroupingKey는 발신자 그룹키이다
// RecipientList는 수신자 리스트이며 최대 1,000명까지 가능하다.
// UserId는 발송구분자이다.
// Stats Id는 통계 ID로 검색 조건에는 포함되지 않는다.
type TextMessage struct {
	TemplateId        string      `json:"templateId" validate:"max=50,required_without=Body"`
	Body              string      `json:"body" validate:"max=255,required_without=TemplateId"`
	SendNo            string      `json:"sendNo" validate:"max=13,required"`
	RequestDate       string      `json:"requestDate" validate:"omitempty,datetime=2006-01-02 15:04"`
	SenderGroupingKey string      `json:"senderGroupingKey" validate:"max=100"`
	RecipientList     []Recipient `json:"recipientList" validate:"max=1000,required"`
	UserId            string      `json:"userId" validate:"max=100"`
	StatsId           string      `json:"statsId" validate:"max=10"`
}

type ResultQuery struct {
	RequestId            string              `json:"requestId" validate:"max=25,required_without_all=StartRequestDate StartCreateDate"`
	StartRequestDate     string              `json:"startRequestDate" validate:"required_without_all=RequestId StartCreateDate,required_with=EndRequestDate,omitempty,datetime=2006-01-02 15:04:05"`
	EndRequestDate       string              `json:"endRequestDate" validate:"required_without_all=RequestId StartCreateDate,required_with=StartRequestDate,omitempty,datetime=2006-01-02 15:04:05"`
	StartCreateDate      string              `json:"startCreateDate" validate:"required_without_all=RequestId StartRequestDate,required_with=EndCreateDate,omitempty,datetime=2006-01-02 15:04:05"`
	EndCreateDate        string              `json:"endCreateDate" validate:"required_without_all=RequestId StartRequestDate,required_with=StartCreateDate,omitempty,datetime=2006-01-02 15:04:05"`
	StartResultDate      string              `json:"startResultDate" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	EndResultDate        string              `json:"endResultDate" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	SendNo               string              `json:"sendNo" validate:"max=13"`
	RecipientNo          string              `json:"recipientNo" validate:"max=20"`
	TemplateId           string              `json:"templateId" validate:"max=50"`
	MsgStatus            MessageStatus       `json:"msgStatus" validate:"omitempty,oneof=0 1 2 3 4 5"`
	ResultCode           ReceiveResult       `json:"resultCode" validate:"omitempty,oneof=MTR1 MTR2"`
	SubResultCode        ReceiveResultDetail `json:"subResultCode" validate:"omitempty,oneof=MTR2_1 MTR2_2 MTR2_3"`
	SenderGroupingKey    string              `json:"senderGroupingKey" validate:"max=100"`
	RecipientGroupingKey string              `json:"recipientGroupingKey" validate:"max=100"`
	PageNum              uint                `json:"pageNum" validate:"omitempty,min=1"`
	PageSize             uint16              `json:"pageSize" validate:"omitempty,min=1,max=1000"`
}
