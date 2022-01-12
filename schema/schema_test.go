package schema

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// Recipient 스키마를 테스트한다.
func TestRecipientSchema(t *testing.T) {
	// 새로운 Validator를 생성한다
	validate := validator.New()

	t.Run("Recipient 객체를 규칙에 맞게 생성했을 때 Validation에 성공해야 한다", func(t *testing.T) {
		recipient := Recipient{
			RecipientNo: "01000000000",
			CountryCode: "82",
			TemplateParameter: map[string]interface{}{
				"hello": "hi",
			},
			RecipientGroupingKey: "RECIPIENT_GROUP",
		}
		err := validate.Struct(recipient)
		// error가 없어야 한다
		assert.Nil(t, err)
	})

	t.Run("Recipient 대신 InternationalRecipientNo를 사용해도 Validation에 성공해야 한다", func(t *testing.T) {
		recipient := Recipient{
			InternationalRecipientNo: "8201000000000",
			TemplateParameter: map[string]interface{}{
				"hello": "hi",
			},
			RecipientGroupingKey: "RECIPIENT_GROUP",
		}
		err := validate.Struct(recipient)
		// error가 없어야 한다
		assert.Nil(t, err)
	})

	t.Run("RecipientNo 혹은 InternationalRecipientNo는 필수이다", func(t *testing.T) {
		recipient := Recipient{
			TemplateParameter: map[string]interface{}{
				"hello": "hi",
			},
			RecipientGroupingKey: "RECIPIENT_GROUP",
		}
		err := validate.Struct(recipient)
		assert.Error(t, err)

		ve := err.(validator.ValidationErrors)
		assert.Equal(t, len(ve), 2)
	})
}

// TextMessage 스키마를 테스트한다.
func TestTextMessageSchema(t *testing.T) {
	validate := validator.New()
	t.Run("TextMessage 객체를 규칙에 맞게 생성했을 때 Validation에 성공해야 한다", func(t *testing.T) {
		// recipient 객체를 생성한다.
		recipient := Recipient{
			RecipientNo: "01000000000",
			TemplateParameter: map[string]interface{}{
				"hello": "hi",
			},
			RecipientGroupingKey: "RECIPIENT_GROUP",
		}
		rerr := validate.Struct(recipient)
		// error가 없어야 한다
		assert.Nil(t, rerr)

		textMessage := TextMessage{
			TemplateId:    "TEST01",
			RecipientList: []Recipient{recipient},
			SendNo:        "0700000000",
		}
		terr := validate.Struct(textMessage)
		assert.Nil(t, terr)
	})
}
