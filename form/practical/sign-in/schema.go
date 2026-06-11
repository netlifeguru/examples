package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func SignInSchema() form.Schema[SignInRequest] {
	f := struct {
		Email    form.StringField[SignInRequest]
		Password form.StringField[SignInRequest]
		Remember form.BoolField[SignInRequest]
	}{
		Email:    form.Str[SignInRequest]("email", func(r *SignInRequest) string { return r.Email }),
		Password: form.Str[SignInRequest]("password", func(r *SignInRequest) string { return r.Password }),
		Remember: form.Bool[SignInRequest]("remember", func(r *SignInRequest) bool { return r.Remember }),
	}

	email := form.Schema[SignInRequest]{
		rules.Required(f.Email),
		rules.Email(f.Email),
	}

	password := form.Schema[SignInRequest]{
		rules.Required(f.Password),
		rules.MinLen(f.Password, 5),
	}

	return form.Rules(
		email,
		password,
	)
}
