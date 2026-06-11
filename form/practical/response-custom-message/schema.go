package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeEmailRequired    = form.Code("email_required")
	CodePasswordRequired = form.Code("password_required")
)

type ResponseCustomMessageRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ResponseCustomMessageSchema() form.Schema[ResponseCustomMessageRequest] {
	ResponseCustomMessageForm := struct {
		Email    form.StringField[ResponseCustomMessageRequest]
		Password form.StringField[ResponseCustomMessageRequest]
	}{
		Email: form.Str[ResponseCustomMessageRequest]("email", func(r *ResponseCustomMessageRequest) string {
			return r.Email
		}),
		Password: form.Str[ResponseCustomMessageRequest]("password", func(r *ResponseCustomMessageRequest) string {
			return r.Password
		}),
	}

	EmailSchema := form.Schema[ResponseCustomMessageRequest]{
		rules.RequiredWithCode(ResponseCustomMessageForm.Email, CodeEmailRequired),
	}

	PasswordSchema := form.Schema[ResponseCustomMessageRequest]{
		rules.RequiredWithCode(ResponseCustomMessageForm.Password, CodePasswordRequired),
	}

	return form.Rules(
		EmailSchema,
		PasswordSchema,
	)
}
