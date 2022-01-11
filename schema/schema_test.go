package schema

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

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
}
