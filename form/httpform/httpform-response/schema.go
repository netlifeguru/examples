package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeEmailRequired    = form.Code("email_required")
	CodePasswordRequired = form.Code("password_required")
	CodePasswordMinLen   = form.Code("password_min_len")
)

type ResponseExampleRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ResponseExampleSchema() form.Schema[ResponseExampleRequest] {
	ResponseExampleForm := struct {
		Email    form.StringField[ResponseExampleRequest]
		Password form.StringField[ResponseExampleRequest]
	}{
		Email: form.Str[ResponseExampleRequest]("email", func(r *ResponseExampleRequest) string {
			return r.Email
		}),
		Password: form.Str[ResponseExampleRequest]("password", func(r *ResponseExampleRequest) string {
			return r.Password
		}),
	}

	EmailSchema := form.Schema[ResponseExampleRequest]{
		rules.RequiredWithCode(ResponseExampleForm.Email, CodeEmailRequired),
		rules.Email(ResponseExampleForm.Email),
	}

	PasswordSchema := form.Schema[ResponseExampleRequest]{
		rules.RequiredWithCode(ResponseExampleForm.Password, CodePasswordRequired),
		rules.MinLenWithCode(ResponseExampleForm.Password, 8, CodePasswordMinLen),
	}

	return form.Rules(
		EmailSchema,
		PasswordSchema,
	)
}
