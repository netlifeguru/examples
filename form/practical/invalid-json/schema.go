package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

type InvalidJSONRequest struct {
	Email string `json:"email"`
}

func InvalidJSONSchema() form.Schema[InvalidJSONRequest] {
	InvalidJSONForm := struct {
		Email form.StringField[InvalidJSONRequest]
	}{
		Email: form.Str[InvalidJSONRequest]("email", func(r *InvalidJSONRequest) string {
			return r.Email
		}),
	}

	return form.Schema[InvalidJSONRequest]{
		rules.Required(InvalidJSONForm.Email),
		rules.Email(InvalidJSONForm.Email),
	}
}
