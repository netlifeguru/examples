package main

import (
	"time"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeRequiredName        = form.Code("required_name")
	CodeRequiredAge         = form.Code("required_age")
	CodeRequiredAccepted    = form.Code("required_accepted")
	CodeRequiredBirthDate   = form.Code("required_birth_date")
	CodeRequiredPermissions = form.Code("required_permissions")
)

type RequiredRulesRequest struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Level       int       `json:"level"`
	Accepted    bool      `json:"accepted"`
	Confirmed   bool      `json:"confirmed"`
	BirthDate   time.Time `json:"birth_date"`
	StartedAt   time.Time `json:"started_at"`
	Permissions []string  `json:"permissions"`
	Roles       []string  `json:"roles"`
}

func RequiredRulesSchema() form.Schema[RequiredRulesRequest] {
	RequiredForm := struct {
		Email       form.StringField[RequiredRulesRequest]
		Name        form.StringField[RequiredRulesRequest]
		Age         form.IntField[RequiredRulesRequest]
		Level       form.IntField[RequiredRulesRequest]
		Accepted    form.BoolField[RequiredRulesRequest]
		Confirmed   form.BoolField[RequiredRulesRequest]
		BirthDate   form.TimeField[RequiredRulesRequest]
		StartedAt   form.TimeField[RequiredRulesRequest]
		Permissions form.SliceStringField[RequiredRulesRequest]
		Roles       form.SliceStringField[RequiredRulesRequest]
	}{
		Email: form.Str[RequiredRulesRequest]("email", func(r *RequiredRulesRequest) string {
			return r.Email
		}),
		Name: form.Str[RequiredRulesRequest]("name", func(r *RequiredRulesRequest) string {
			return r.Name
		}),
		Age: form.Int[RequiredRulesRequest]("age", func(r *RequiredRulesRequest) int {
			return r.Age
		}),
		Level: form.Int[RequiredRulesRequest]("level", func(r *RequiredRulesRequest) int {
			return r.Level
		}),
		Accepted: form.Bool[RequiredRulesRequest]("accepted", func(r *RequiredRulesRequest) bool {
			return r.Accepted
		}),
		Confirmed: form.Bool[RequiredRulesRequest]("confirmed", func(r *RequiredRulesRequest) bool {
			return r.Confirmed
		}),
		BirthDate: form.Time[RequiredRulesRequest]("birth_date", func(r *RequiredRulesRequest) time.Time {
			return r.BirthDate
		}),
		StartedAt: form.Time[RequiredRulesRequest]("started_at", func(r *RequiredRulesRequest) time.Time {
			return r.StartedAt
		}),
		Permissions: form.SliceStr[RequiredRulesRequest]("permissions", func(r *RequiredRulesRequest) []string {
			return r.Permissions
		}),
		Roles: form.SliceStr[RequiredRulesRequest]("roles", func(r *RequiredRulesRequest) []string {
			return r.Roles
		}),
	}

	return form.Schema[RequiredRulesRequest]{
		// Required validates that a string value is not empty.
		rules.Required(RequiredForm.Email),

		// RequiredWithCode validates string value and returns a custom error code.
		rules.RequiredWithCode(RequiredForm.Name, CodeRequiredName),

		// RequiredInt validates that an int value is not zero.
		rules.RequiredInt(RequiredForm.Age),

		// RequiredIntWithCode validates int value and returns a custom error code.
		rules.RequiredIntWithCode(RequiredForm.Level, CodeRequiredAge),

		// RequiredBool validates that a bool value is true.
		rules.RequiredBool(RequiredForm.Accepted),

		// RequiredBoolWithCode validates bool value and returns a custom error code.
		rules.RequiredBoolWithCode(RequiredForm.Confirmed, CodeRequiredAccepted),

		// RequiredTime validates that a time value is not zero.
		rules.RequiredTime(RequiredForm.BirthDate),

		// RequiredTimeWithCode validates time value and returns a custom error code.
		rules.RequiredTimeWithCode(RequiredForm.StartedAt, CodeRequiredBirthDate),

		// RequiredSliceStr validates that a string slice is not empty.
		rules.RequiredSliceStr(RequiredForm.Permissions),

		// RequiredSliceStrWithCode validates string slice and returns a custom error code.
		rules.RequiredSliceStrWithCode(RequiredForm.Roles, CodeRequiredPermissions),
	}
}
