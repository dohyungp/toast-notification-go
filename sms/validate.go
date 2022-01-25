package sms

import "github.com/go-playground/validator/v10"

func validateTextMessageMaxLength(sl validator.StructLevel) {
	message := sl.Current().Interface().(TextMessage)
	var bodyMaxLength int
	var titleRequired bool

	switch message.Type {
	case "sms":
		titleRequired = false
		bodyMaxLength = 256
	case "mms":
		titleRequired = true
		bodyMaxLength = 4000
	}

	if message.TemplateId == "" {
		if titleRequired && message.Title == "" {
			sl.ReportError(message.Title, "Title", "Title", "mmstitle", "")
		}

		if len(message.Body) > bodyMaxLength {
			sl.ReportError(message.Body, "Body", "Body", "bodylength", "")
		}
	}
}

func validateResultQuery(sl validator.StructLevel) {
	resultQuery := sl.Current().Interface().(ResultQuery)

	switch {
	case resultQuery.RequestId != "":
		break
	case resultQuery.StartCreateDate != "":
		if resultQuery.EndCreateDate == "" {
			sl.ReportError(resultQuery.EndCreateDate, "EndCreateDate", "EndCreateDate", "endcreatedate", "")
		}
	case resultQuery.StartRequestDate != "":
		if resultQuery.EndRequestDate == "" {
			sl.ReportError(resultQuery.EndRequestDate, "EndRequestDate", "EndRequestDate", "endrequestdate", "")
		}
	case resultQuery.StartResultDate != "":
		if resultQuery.EndResultDate == "" {
			sl.ReportError(resultQuery.EndResultDate, "EndResultDate", "EndResultDate", "endresultdate", "")
		}
	default:
		// FIXME: Exception을 좀더 친절하게 수정이 필요하다.
		sl.ReportError(resultQuery.RequestId, "RequestId", "RequestId", "requestid", "")
	}

}

// Schema의 구조체 레벨 Validaton을 수행하는 Validator구조체 생성자이다.
func NewValidate() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(validateTextMessageMaxLength, TextMessage{})
	validate.RegisterStructValidation(validateResultQuery, ResultQuery{})
	return validate
}
