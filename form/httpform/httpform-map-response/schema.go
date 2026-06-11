package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

type MapResponseRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func MapResponseSchema() form.Schema[MapResponseRequest] {
	MapResponseForm := struct {
		Email    form.StringField[MapResponseRequest]
		Password form.StringField[MapResponseRequest]
	}{
		Email: form.Str[MapResponseRequest]("email", func(r *MapResponseRequest) string {
			return r.Email
		}),
		Password: form.Str[MapResponseRequest]("password", func(r *MapResponseRequest) string {
			return r.Password
		}),
	}

	EmailSchema := form.Schema[MapResponseRequest]{
		rules.Required(MapResponseForm.Email),
		rules.Email(MapResponseForm.Email),
	}

	PasswordSchema := form.Schema[MapResponseRequest]{
		rules.Required(MapResponseForm.Password),
		rules.MinLen(MapResponseForm.Password, 8),
	}

	return form.Rules(
		EmailSchema,
		PasswordSchema,
	)
}
