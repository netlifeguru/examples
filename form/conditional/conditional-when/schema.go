package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/conditional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodePasswordRequired        = form.Code("password_required")
	CodePasswordMinLen          = form.Code("password_min_len")
	CodePasswordConfirmRequired = form.Code("password_confirm_required")

	CodeCompanyNameRequired = form.Code("company_name_required")
	CodeVatNumberRequired   = form.Code("vat_number_required")
)

type ConditionalWhenRequest struct {
	ChangePassword  bool   `json:"change_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`

	AccountType string `json:"account_type"`
	CompanyName string `json:"company_name"`
	VatNumber   string `json:"vat_number"`
}

func ConditionalWhenSchema() form.Schema[ConditionalWhenRequest] {
	ConditionalWhenForm := struct {
		ChangePassword  form.BoolField[ConditionalWhenRequest]
		Password        form.StringField[ConditionalWhenRequest]
		ConfirmPassword form.StringField[ConditionalWhenRequest]

		AccountType form.StringField[ConditionalWhenRequest]
		CompanyName form.StringField[ConditionalWhenRequest]
		VatNumber   form.StringField[ConditionalWhenRequest]
	}{
		ChangePassword: form.Bool[ConditionalWhenRequest]("change_password", func(r *ConditionalWhenRequest) bool {
			return r.ChangePassword
		}),
		Password: form.Str[ConditionalWhenRequest]("password", func(r *ConditionalWhenRequest) string {
			return r.Password
		}),
		ConfirmPassword: form.Str[ConditionalWhenRequest]("confirm_password", func(r *ConditionalWhenRequest) string {
			return r.ConfirmPassword
		}),
		AccountType: form.Str[ConditionalWhenRequest]("account_type", func(r *ConditionalWhenRequest) string {
			return r.AccountType
		}),
		CompanyName: form.Str[ConditionalWhenRequest]("company_name", func(r *ConditionalWhenRequest) string {
			return r.CompanyName
		}),
		VatNumber: form.Str[ConditionalWhenRequest]("vat_number", func(r *ConditionalWhenRequest) string {
			return r.VatNumber
		}),
	}

	return form.Schema[ConditionalWhenRequest]{
		// Run password validation only when change_password is true.
		conditional.When(
			func(r *ConditionalWhenRequest) bool {
				return r.ChangePassword
			},
			rules.RequiredWithCode(ConditionalWhenForm.Password, CodePasswordRequired),
			rules.MinLenWithCode(ConditionalWhenForm.Password, 8, CodePasswordMinLen),
			rules.RequiredWithCode(ConditionalWhenForm.ConfirmPassword, CodePasswordConfirmRequired),
		),

		// Run company validation only when account_type is "company".
		conditional.When(
			func(r *ConditionalWhenRequest) bool {
				return r.AccountType == "company"
			},
			rules.RequiredWithCode(ConditionalWhenForm.CompanyName, CodeCompanyNameRequired),
			rules.RequiredWithCode(ConditionalWhenForm.VatNumber, CodeVatNumberRequired),
		),
	}
}
