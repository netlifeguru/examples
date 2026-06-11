package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeEmailRequired    = form.Code("email_required")
	CodeEmailInvalid     = form.Code("email_invalid")
	CodePasswordRequired = form.Code("password_required")
	CodePasswordMinLen   = form.Code("password_min_len")
	CodeTermsRequired    = form.Code("terms_required")
)

type RegistrationRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Age             int    `json:"age"`
	AcceptTerms     bool   `json:"accept_terms"`
}

func RegistrationSchema() form.Schema[RegistrationRequest] {
	RegistrationForm := struct {
		Name            form.StringField[RegistrationRequest]
		Email           form.StringField[RegistrationRequest]
		Password        form.StringField[RegistrationRequest]
		PasswordConfirm form.StringField[RegistrationRequest]
		Age             form.IntField[RegistrationRequest]
		AcceptTerms     form.BoolField[RegistrationRequest]
	}{
		Name: form.Str[RegistrationRequest]("name", func(r *RegistrationRequest) string {
			return r.Name
		}),
		Email: form.Str[RegistrationRequest]("email", func(r *RegistrationRequest) string {
			return r.Email
		}),
		Password: form.Str[RegistrationRequest]("password", func(r *RegistrationRequest) string {
			return r.Password
		}),
		PasswordConfirm: form.Str[RegistrationRequest]("password_confirm", func(r *RegistrationRequest) string {
			return r.PasswordConfirm
		}),
		Age: form.Int[RegistrationRequest]("age", func(r *RegistrationRequest) int {
			return r.Age
		}),
		AcceptTerms: form.Bool[RegistrationRequest]("accept_terms", func(r *RegistrationRequest) bool {
			return r.AcceptTerms
		}),
	}

	NameSchema := form.Schema[RegistrationRequest]{
		rules.Required(RegistrationForm.Name),
		rules.MinLen(RegistrationForm.Name, 2),
	}

	EmailSchema := form.Schema[RegistrationRequest]{
		rules.RequiredWithCode(RegistrationForm.Email, CodeEmailRequired),
		rules.EmailWithCode(RegistrationForm.Email, CodeEmailInvalid),
	}

	PasswordSchema := form.Schema[RegistrationRequest]{
		rules.RequiredWithCode(RegistrationForm.Password, CodePasswordRequired),
		rules.MinLenWithCode(RegistrationForm.Password, 8, CodePasswordMinLen),
		rules.Required(RegistrationForm.PasswordConfirm),
		rules.MinLen(RegistrationForm.PasswordConfirm, 8),
	}

	AgeSchema := form.Schema[RegistrationRequest]{
		rules.RequiredInt(RegistrationForm.Age),
		rules.Min(RegistrationForm.Age, 18),
	}

	TermsSchema := form.Schema[RegistrationRequest]{
		rules.IsTrueWithCode(RegistrationForm.AcceptTerms, CodeTermsRequired),
	}

	return form.Rules(
		NameSchema,
		EmailSchema,
		PasswordSchema,
		AgeSchema,
		TermsSchema,
	)
}
