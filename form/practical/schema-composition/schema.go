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
	CodeNameRequired     = form.Code("name_required")
	CodeAgeMin           = form.Code("age_min")
)

type SchemaCompositionRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func SchemaCompositionSchema() form.Schema[SchemaCompositionRequest] {
	SchemaCompositionForm := struct {
		Name     form.StringField[SchemaCompositionRequest]
		Email    form.StringField[SchemaCompositionRequest]
		Password form.StringField[SchemaCompositionRequest]
		Age      form.IntField[SchemaCompositionRequest]
	}{
		Name: form.Str[SchemaCompositionRequest]("name", func(r *SchemaCompositionRequest) string {
			return r.Name
		}),
		Email: form.Str[SchemaCompositionRequest]("email", func(r *SchemaCompositionRequest) string {
			return r.Email
		}),
		Password: form.Str[SchemaCompositionRequest]("password", func(r *SchemaCompositionRequest) string {
			return r.Password
		}),
		Age: form.Int[SchemaCompositionRequest]("age", func(r *SchemaCompositionRequest) int {
			return r.Age
		}),
	}

	NameSchema := form.Schema[SchemaCompositionRequest]{
		rules.RequiredWithCode(SchemaCompositionForm.Name, CodeNameRequired),
		rules.MinLen(SchemaCompositionForm.Name, 2),
	}

	EmailSchema := form.Schema[SchemaCompositionRequest]{
		rules.RequiredWithCode(SchemaCompositionForm.Email, CodeEmailRequired),
		rules.EmailWithCode(SchemaCompositionForm.Email, CodeEmailInvalid),
	}

	PasswordSchema := form.Schema[SchemaCompositionRequest]{
		rules.RequiredWithCode(SchemaCompositionForm.Password, CodePasswordRequired),
		rules.MinLenWithCode(SchemaCompositionForm.Password, 8, CodePasswordMinLen),
	}

	AgeSchema := form.Schema[SchemaCompositionRequest]{
		rules.RequiredInt(SchemaCompositionForm.Age),
		rules.MinWithCode(SchemaCompositionForm.Age, 18, CodeAgeMin),
	}

	return form.Rules(
		NameSchema,
		EmailSchema,
		PasswordSchema,
		AgeSchema,
	)
}
