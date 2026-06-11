package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/conditional"
)

const (
	CodeCompanyNameRequired = form.Code("company_name_required")
	CodeVatNumberRequired   = form.Code("vat_number_required")
	CodeAdminNoteForbidden  = form.Code("admin_note_forbidden")
	CodeInternalNoteBlocked = form.Code("internal_note_blocked")
	CodeEmailRequired       = form.Code("email_required")
	CodeContactRequired     = form.Code("contact_required")
	CodePhoneRequired       = form.Code("phone_required")
	CodeBackupRequired      = form.Code("backup_required")
)

type ConditionalRulesRequest struct {
	AccountType string `json:"account_type"`
	CompanyName string `json:"company_name"`

	Country   string `json:"country"`
	VatNumber string `json:"vat_number"`

	Role      string `json:"role"`
	AdminNote string `json:"admin_note"`

	UserType     string `json:"user_type"`
	InternalNote string `json:"internal_note"`

	Phone string `json:"phone"`
	Email string `json:"email"`

	Address     string `json:"address"`
	City        string `json:"city"`
	ContactName string `json:"contact_name"`

	BackupEmail string `json:"backup_email"`
	BackupPhone string `json:"backup_phone"`
}

func ConditionalRulesSchema() form.Schema[ConditionalRulesRequest] {
	ConditionalRulesForm := struct {
		AccountType form.StringField[ConditionalRulesRequest]
		CompanyName form.StringField[ConditionalRulesRequest]

		Country   form.StringField[ConditionalRulesRequest]
		VatNumber form.StringField[ConditionalRulesRequest]

		Role      form.StringField[ConditionalRulesRequest]
		AdminNote form.StringField[ConditionalRulesRequest]

		UserType     form.StringField[ConditionalRulesRequest]
		InternalNote form.StringField[ConditionalRulesRequest]

		Phone form.StringField[ConditionalRulesRequest]
		Email form.StringField[ConditionalRulesRequest]

		Address     form.StringField[ConditionalRulesRequest]
		City        form.StringField[ConditionalRulesRequest]
		ContactName form.StringField[ConditionalRulesRequest]

		BackupEmail form.StringField[ConditionalRulesRequest]
		BackupPhone form.StringField[ConditionalRulesRequest]
	}{
		AccountType: form.Str[ConditionalRulesRequest]("account_type", func(r *ConditionalRulesRequest) string {
			return r.AccountType
		}),
		CompanyName: form.Str[ConditionalRulesRequest]("company_name", func(r *ConditionalRulesRequest) string {
			return r.CompanyName
		}),
		Country: form.Str[ConditionalRulesRequest]("country", func(r *ConditionalRulesRequest) string {
			return r.Country
		}),
		VatNumber: form.Str[ConditionalRulesRequest]("vat_number", func(r *ConditionalRulesRequest) string {
			return r.VatNumber
		}),
		Role: form.Str[ConditionalRulesRequest]("role", func(r *ConditionalRulesRequest) string {
			return r.Role
		}),
		AdminNote: form.Str[ConditionalRulesRequest]("admin_note", func(r *ConditionalRulesRequest) string {
			return r.AdminNote
		}),
		UserType: form.Str[ConditionalRulesRequest]("user_type", func(r *ConditionalRulesRequest) string {
			return r.UserType
		}),
		InternalNote: form.Str[ConditionalRulesRequest]("internal_note", func(r *ConditionalRulesRequest) string {
			return r.InternalNote
		}),
		Phone: form.Str[ConditionalRulesRequest]("phone", func(r *ConditionalRulesRequest) string {
			return r.Phone
		}),
		Email: form.Str[ConditionalRulesRequest]("email", func(r *ConditionalRulesRequest) string {
			return r.Email
		}),
		Address: form.Str[ConditionalRulesRequest]("address", func(r *ConditionalRulesRequest) string {
			return r.Address
		}),
		City: form.Str[ConditionalRulesRequest]("city", func(r *ConditionalRulesRequest) string {
			return r.City
		}),
		ContactName: form.Str[ConditionalRulesRequest]("contact_name", func(r *ConditionalRulesRequest) string {
			return r.ContactName
		}),
		BackupEmail: form.Str[ConditionalRulesRequest]("backup_email", func(r *ConditionalRulesRequest) string {
			return r.BackupEmail
		}),
		BackupPhone: form.Str[ConditionalRulesRequest]("backup_phone", func(r *ConditionalRulesRequest) string {
			return r.BackupPhone
		}),
	}

	return form.Schema[ConditionalRulesRequest]{
		// RequiredIfStr requires company_name when account_type is "company".
		conditional.RequiredIfStr(
			ConditionalRulesForm.CompanyName,
			func(r *ConditionalRulesRequest) bool {
				return r.AccountType == "company"
			},
		),

		// RequiredIfStrWithCode requires vat_number when country is "SK".
		conditional.RequiredIfStrWithCode(
			ConditionalRulesForm.VatNumber,
			func(r *ConditionalRulesRequest) bool {
				return r.Country == "SK"
			},
			CodeVatNumberRequired,
		),

		// ProhibitedIfStr forbids admin_note when role is not "admin".
		conditional.ProhibitedIfStr(
			ConditionalRulesForm.AdminNote,
			func(r *ConditionalRulesRequest) bool {
				return r.Role != "admin"
			},
		),

		// ProhibitedIfStrWithCode forbids internal_note when user_type is "external".
		conditional.ProhibitedIfStrWithCode(
			ConditionalRulesForm.InternalNote,
			func(r *ConditionalRulesRequest) bool {
				return r.UserType == "external"
			},
			CodeInternalNoteBlocked,
		),

		// RequiredWithAnyStr requires email when phone is filled.
		conditional.RequiredWithAnyStr(
			ConditionalRulesForm.Email,
			ConditionalRulesForm.Phone,
		),

		// RequiredWithAnyStrWithCode requires contact_name when address or city is filled.
		conditional.RequiredWithAnyStrWithCode(
			ConditionalRulesForm.ContactName,
			CodeContactRequired,
			ConditionalRulesForm.Address,
			ConditionalRulesForm.City,
		),

		// RequiredWithoutAnyStr requires phone when email is blank.
		conditional.RequiredWithoutAnyStr(
			ConditionalRulesForm.Phone,
			ConditionalRulesForm.Email,
		),

		// RequiredWithoutAnyStrWithCode requires backup_email when backup_phone is blank.
		conditional.RequiredWithoutAnyStrWithCode(
			ConditionalRulesForm.BackupEmail,
			CodeBackupRequired,
			ConditionalRulesForm.BackupPhone,
		),
	}
}
