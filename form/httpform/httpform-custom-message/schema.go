package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeEmailRequired    = form.Code("email_required")
	CodePasswordRequired = form.Code("password_required")
)

type CustomMessageRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CustomMessageSchema() form.Schema[CustomMessageRequest] {
	CustomMessageForm := struct {
		Email    form.StringField[CustomMessageRequest]
		Password form.StringField[CustomMessageRequest]
	}{
		Email: form.Str[CustomMessageRequest]("email", func(r *CustomMessageRequest) string {
			return r.Email
		}),
		Password: form.Str[CustomMessageRequest]("password", func(r *CustomMessageRequest) string {
			return r.Password
		}),
	}

	EmailSchema := form.Schema[CustomMessageRequest]{
		rules.RequiredWithCode(CustomMessageForm.Email, CodeEmailRequired),
		rules.Email(CustomMessageForm.Email),
	}

	PasswordSchema := form.Schema[CustomMessageRequest]{
		rules.RequiredWithCode(CustomMessageForm.Password, CodePasswordRequired),
	}

	return form.Rules(
		EmailSchema,
		PasswordSchema,
	)
}
