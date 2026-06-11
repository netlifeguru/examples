package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodePasswordsMismatch = form.Code("passwords_mismatch")
	CodeInvalidRange      = form.Code("invalid_range")
)

type CompareRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`

	MinPrice int `json:"min_price"`
	MaxPrice int `json:"max_price"`
}

func CompareSchema() form.Schema[CompareRequest] {
	CompareForm := struct {
		Password        form.StringField[CompareRequest]
		ConfirmPassword form.StringField[CompareRequest]

		MinPrice form.IntField[CompareRequest]
		MaxPrice form.IntField[CompareRequest]
	}{
		Password: form.Str[CompareRequest]("password", func(r *CompareRequest) string {
			return r.Password
		}),
		ConfirmPassword: form.Str[CompareRequest]("confirm_password", func(r *CompareRequest) string {
			return r.ConfirmPassword
		}),

		MinPrice: form.Int[CompareRequest]("min_price", func(r *CompareRequest) int {
			return r.MinPrice
		}),
		MaxPrice: form.Int[CompareRequest]("max_price", func(r *CompareRequest) int {
			return r.MaxPrice
		}),
	}

	return form.Schema[CompareRequest]{
		rules.CompareWithCode(
			CompareForm.Password.Field,
			CompareForm.ConfirmPassword.Field,
			rules.OpEQ,
			CodePasswordsMismatch,
		),

		rules.CompareWithCode(
			CompareForm.MinPrice.Field,
			CompareForm.MaxPrice.Field,
			rules.OpLTE,
			CodeInvalidRange,
		),
	}
}
