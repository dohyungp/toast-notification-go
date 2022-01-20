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

func NewValidate() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(validateTextMessageMaxLength, TextMessage{})
	return validate
}
