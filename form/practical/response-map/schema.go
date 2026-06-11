package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

type ResponseMapRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ResponseMapSchema() form.Schema[ResponseMapRequest] {
	ResponseMapForm := struct {
		Email    form.StringField[ResponseMapRequest]
		Password form.StringField[ResponseMapRequest]
	}{
		Email: form.Str[ResponseMapRequest]("email", func(r *ResponseMapRequest) string {
			return r.Email
		}),
		Password: form.Str[ResponseMapRequest]("password", func(r *ResponseMapRequest) string {
			return r.Password
		}),
	}

	EmailSchema := form.Schema[ResponseMapRequest]{
		rules.Required(ResponseMapForm.Email),
		rules.Email(ResponseMapForm.Email),
	}

	PasswordSchema := form.Schema[ResponseMapRequest]{
		rules.Required(ResponseMapForm.Password),
		rules.MinLen(ResponseMapForm.Password, 8),
	}

	return form.Rules(
		EmailSchema,
		PasswordSchema,
	)
}
