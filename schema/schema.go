package schema

// Toast SMS API 수신인이다
// RecipientNo는 전화번호이며 필수이다
// CountryCode는 국가코드이며 없는 경우 자동으로 82가 된다
// InternationalRecipentNo는 국가코드를 포함한 전화번호이며 recipentNo가 있는 경우 무시된다
// TemplateParameter Toast에서 동적으로 문자 템플릿을 만들때 사용하며 Map 타입이다
// RecipientGroupingKey는 수신자 그룹키이다
type Recipient struct {
	RecipientNo              string                 `json:"recipientNo" validate:"max=50"`
	CountryCode              string                 `json:"countryCode" validate:"max=8"`
	InternationalRecipientNo string                 `json:"internationalRecipientNo" validate:"e164,max=20"`
	TemplateParameter        map[string]interface{} `json:"templateParameter,omitempty" validate:"omitempty,max=20"`
	RecipientGroupingKey     string                 `json:"recipientGroupingKey,omitempty" validate:"max=100"`
}
