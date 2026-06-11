package main

import (
	"strings"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/conditional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeCompanyNameRequired = form.Code("company_name_required")
	CodeAdminNoteForbidden  = form.Code("admin_note_forbidden")
	CodeEmailRequired       = form.Code("email_required")
	CodePhoneRequired       = form.Code("phone_required")
	CodeContactNameRequired = form.Code("contact_name_required")
	CodeVatNumberRequired   = form.Code("vat_number_required")
	CodePasswordMinLen      = form.Code("password_min_len")
)

type ConditionalRulesRequest struct {
	AccountType string `json:"account_type"`
	CompanyName string `json:"company_name"`

	Role      string `json:"role"`
	AdminNote string `json:"admin_note"`

	Email string `json:"email"`
	Phone string `json:"phone"`

	ContactName string `json:"contact_name"`
	Address     string `json:"address"`
	City        string `json:"city"`

	Country   string `json:"country"`
	VatNumber string `json:"vat_number"`

	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ConditionalRulesSchema() form.Schema[ConditionalRulesRequest] {
	ConditionalForm := struct {
		AccountType     form.StringField[ConditionalRulesRequest]
		CompanyName     form.StringField[ConditionalRulesRequest]
		Role            form.StringField[ConditionalRulesRequest]
		AdminNote       form.StringField[ConditionalRulesRequest]
		Email           form.StringField[ConditionalRulesRequest]
		Phone           form.StringField[ConditionalRulesRequest]
		ContactName     form.StringField[ConditionalRulesRequest]
		Address         form.StringField[ConditionalRulesRequest]
		City            form.StringField[ConditionalRulesRequest]
		Country         form.StringField[ConditionalRulesRequest]
		VatNumber       form.StringField[ConditionalRulesRequest]
		Password        form.StringField[ConditionalRulesRequest]
		ConfirmPassword form.StringField[ConditionalRulesRequest]
	}{
		AccountType: form.Str[ConditionalRulesRequest]("account_type", func(r *ConditionalRulesRequest) string {
			return r.AccountType
		}),
		CompanyName: form.Str[ConditionalRulesRequest]("company_name", func(r *ConditionalRulesRequest) string {
			return r.CompanyName
		}),
		Role: form.Str[ConditionalRulesRequest]("role", func(r *ConditionalRulesRequest) string {
			return r.Role
		}),
		AdminNote: form.Str[ConditionalRulesRequest]("admin_note", func(r *ConditionalRulesRequest) string {
			return r.AdminNote
		}),
		Email: form.Str[ConditionalRulesRequest]("email", func(r *ConditionalRulesRequest) string {
			return r.Email
		}),
		Phone: form.Str[ConditionalRulesRequest]("phone", func(r *ConditionalRulesRequest) string {
			return r.Phone
		}),
		ContactName: form.Str[ConditionalRulesRequest]("contact_name", func(r *ConditionalRulesRequest) string {
			return r.ContactName
		}),
		Address: form.Str[ConditionalRulesRequest]("address", func(r *ConditionalRulesRequest) string {
			return r.Address
		}),
		City: form.Str[ConditionalRulesRequest]("city", func(r *ConditionalRulesRequest) string {
			return r.City
		}),
		Country: form.Str[ConditionalRulesRequest]("country", func(r *ConditionalRulesRequest) string {
			return r.Country
		}),
		VatNumber: form.Str[ConditionalRulesRequest]("vat_number", func(r *ConditionalRulesRequest) string {
			return r.VatNumber
		}),
		Password: form.Str[ConditionalRulesRequest]("password", func(r *ConditionalRulesRequest) string {
			return r.Password
		}),
		ConfirmPassword: form.Str[ConditionalRulesRequest]("confirm_password", func(r *ConditionalRulesRequest) string {
			return r.ConfirmPassword
		}),
	}

	return form.Schema[ConditionalRulesRequest]{
		// RequiredIfStr requires company_name when account_type is "company".
		conditional.RequiredIfStr(
			ConditionalForm.CompanyName,
			func(r *ConditionalRulesRequest) bool {
				return r.AccountType == "company"
			},
		),

		// RequiredIfStrWithCode requires vat_number when country is "SK".
		conditional.RequiredIfStrWithCode(
			ConditionalForm.VatNumber,
			func(r *ConditionalRulesRequest) bool {
				return r.Country == "SK"
			},
			CodeVatNumberRequired,
		),

		// ProhibitedIfStr forbids admin_note when role is not "admin".
		conditional.ProhibitedIfStr(
			ConditionalForm.AdminNote,
			func(r *ConditionalRulesRequest) bool {
				return r.Role != "admin"
			},
		),

		// ProhibitedIfStrWithCode forbids company_name when account_type is "personal".
		conditional.ProhibitedIfStrWithCode(
			ConditionalForm.CompanyName,
			func(r *ConditionalRulesRequest) bool {
				return r.AccountType == "personal"
			},
			CodeAdminNoteForbidden,
		),

		// RequiredWithAnyStr requires email when phone is filled.
		conditional.RequiredWithAnyStr(
			ConditionalForm.Email,
			ConditionalForm.Phone,
		),

		// RequiredWithAnyStrWithCode requires contact_name when address or city is filled.
		conditional.RequiredWithAnyStrWithCode(
			ConditionalForm.ContactName,
			CodeContactNameRequired,
			ConditionalForm.Address,
			ConditionalForm.City,
		),

		// RequiredWithoutAnyStr requires phone when email is blank.
		conditional.RequiredWithoutAnyStr(
			ConditionalForm.Phone,
			ConditionalForm.Email,
		),

		// RequiredWithoutAnyStrWithCode requires email when phone is blank.
		conditional.RequiredWithoutAnyStrWithCode(
			ConditionalForm.Email,
			CodeEmailRequired,
			ConditionalForm.Phone,
		),

		// When applies nested rules only when condition is true.
		conditional.When(
			func(r *ConditionalRulesRequest) bool {
				return strings.TrimSpace(r.Password) != ""
			},
			rules.MinLenWithCode(ConditionalForm.Password, 8, CodePasswordMinLen),
			rules.Required(ConditionalForm.ConfirmPassword),
		),
	}
}
