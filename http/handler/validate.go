package handler

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateEmail(v string) error {
	return validation.Validate(v,
		validation.Required.Error("Emailは必須です。"),
		is.Email.Error("Emailのフォーマットが不正です。"),
	)
}

func ValidatePassword(v string) error {
	return validation.Validate(v,
		validation.Required.Error("パスワードは必須です"),
		validation.Length(4, 100).Error("パスワードは4~100文字です"),
	)
}

func ValidateName(v string) error {
	return validation.Validate(v,
		validation.Required.Error("Nameは必須です"),
	)
}