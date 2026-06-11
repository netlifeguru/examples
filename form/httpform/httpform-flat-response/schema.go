package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

type FlatResponseRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FlatResponseSchema() form.Schema[FlatResponseRequest] {
	FlatResponseForm := struct {
		Email    form.StringField[FlatResponseRequest]
		Password form.StringField[FlatResponseRequest]
	}{
		Email: form.Str[FlatResponseRequest]("email", func(r *FlatResponseRequest) string {
			return r.Email
		}),
		Password: form.Str[FlatResponseRequest]("password", func(r *FlatResponseRequest) string {
			return r.Password
		}),
	}

	EmailSchema := form.Schema[FlatResponseRequest]{
		rules.Required(FlatResponseForm.Email),
		rules.Email(FlatResponseForm.Email),
	}

	PasswordSchema := form.Schema[FlatResponseRequest]{
		rules.Required(FlatResponseForm.Password),
		rules.MinLen(FlatResponseForm.Password, 8),
	}

	return form.Rules(
		EmailSchema,
		PasswordSchema,
	)
}
