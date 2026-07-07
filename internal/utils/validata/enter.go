package validata

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zhtranslations.RegisterDefaultTranslations(v, trans)
	}
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
}

func ValidateErr(err error) string {
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		return err.Error()
	}

	list := make([]string, 0, len(errs))
	for _, e := range errs {
		list = append(list, e.Translate(trans))
	}

	return strings.Join(list, ";")
}
