package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeInvalidStatus   = form.Code("invalid_status")
	CodeInvalidPlan     = form.Code("invalid_plan")
	CodeInvalidPriority = form.Code("invalid_priority")
)

type GenericRulesRequest struct {
	Role     string `json:"role"`
	Status   string `json:"status"`
	Plan     string `json:"plan"`
	Level    int    `json:"level"`
	Priority int    `json:"priority"`
	Active   bool   `json:"active"`
}

func GenericRulesSchema() form.Schema[GenericRulesRequest] {
	GenericForm := struct {
		Role     form.StringField[GenericRulesRequest]
		Status   form.StringField[GenericRulesRequest]
		Plan     form.StringField[GenericRulesRequest]
		Level    form.IntField[GenericRulesRequest]
		Priority form.IntField[GenericRulesRequest]
		Active   form.BoolField[GenericRulesRequest]
	}{
		Role: form.Str[GenericRulesRequest]("role", func(r *GenericRulesRequest) string {
			return r.Role
		}),
		Status: form.Str[GenericRulesRequest]("status", func(r *GenericRulesRequest) string {
			return r.Status
		}),
		Plan: form.Str[GenericRulesRequest]("plan", func(r *GenericRulesRequest) string {
			return r.Plan
		}),
		Level: form.Int[GenericRulesRequest]("level", func(r *GenericRulesRequest) int {
			return r.Level
		}),
		Priority: form.Int[GenericRulesRequest]("priority", func(r *GenericRulesRequest) int {
			return r.Priority
		}),
		Active: form.Bool[GenericRulesRequest]("active", func(r *GenericRulesRequest) bool {
			return r.Active
		}),
	}

	return form.Schema[GenericRulesRequest]{
		rules.OneOf[GenericRulesRequest, string](
			GenericForm.Role.Field,
			"admin",
			"editor",
			"viewer",
		),

		rules.OneOfWithCode[GenericRulesRequest, string](
			GenericForm.Status.Field,
			CodeInvalidStatus,
			"draft",
			"published",
			"archived",
		),

		rules.OneOfWithCode[GenericRulesRequest, string](
			GenericForm.Plan.Field,
			CodeInvalidPlan,
			"free",
			"pro",
			"enterprise",
		),

		rules.OneOf[GenericRulesRequest, int](
			GenericForm.Level.Field,
			1,
			2,
			3,
		),

		rules.OneOfWithCode[GenericRulesRequest, int](
			GenericForm.Priority.Field,
			CodeInvalidPriority,
			10,
			20,
			30,
		),

		rules.OneOf[GenericRulesRequest, bool](
			GenericForm.Active.Field,
			true,
		),
	}
}
