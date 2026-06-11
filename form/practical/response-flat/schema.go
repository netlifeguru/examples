package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

type ResponseFlatRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ResponseFlatSchema() form.Schema[ResponseFlatRequest] {
	ResponseFlatForm := struct {
		Email    form.StringField[ResponseFlatRequest]
		Password form.StringField[ResponseFlatRequest]
	}{
		Email: form.Str[ResponseFlatRequest]("email", func(r *ResponseFlatRequest) string {
			return r.Email
		}),
		Password: form.Str[ResponseFlatRequest]("password", func(r *ResponseFlatRequest) string {
			return r.Password
		}),
	}

	EmailSchema := form.Schema[ResponseFlatRequest]{
		rules.Required(ResponseFlatForm.Email),
		rules.Email(ResponseFlatForm.Email),
	}

	PasswordSchema := form.Schema[ResponseFlatRequest]{
		rules.Required(ResponseFlatForm.Password),
		rules.MinLen(ResponseFlatForm.Password, 8),
	}

	return form.Rules(
		EmailSchema,
		PasswordSchema,
	)
}
