package web

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func registerTranslator() (*validator.Validate, error) {
	validate := validator.New()

	english := en.New()
	translator = ut.New(english, english)
	translatorEng, ok := translator.GetTranslator("en")
	if !ok {
		return nil, errors.New("translator not found")
	}
	err := enTranslation.RegisterDefaultTranslations(validate, translatorEng)
	if err != nil {
		return nil, errors.New("translator not found")
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate, nil
}

var (
	validate   *validator.Validate
	translator *ut.UniversalTranslator
)

func Validate() *validator.Validate {
	return validate
}
func init() {
	var err error
	validate, err = registerTranslator()
	if err != nil {
		panic(err)
	}
}
