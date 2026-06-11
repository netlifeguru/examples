package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/conditional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeCompanyNameRequired  = form.Code("company_name_required")
	CodeVatNumberRequired    = form.Code("vat_number_required")
	CodeBillingEmailInvalid  = form.Code("billing_email_invalid")
	CodeBillingEmailRequired = form.Code("billing_email_required")
)

type CompanyBillingRequest struct {
	AccountType  string `json:"account_type"`
	CompanyName  string `json:"company_name"`
	Country      string `json:"country"`
	VatNumber    string `json:"vat_number"`
	BillingEmail string `json:"billing_email"`
}

func CompanyBillingSchema() form.Schema[CompanyBillingRequest] {
	CompanyBillingForm := struct {
		AccountType  form.StringField[CompanyBillingRequest]
		CompanyName  form.StringField[CompanyBillingRequest]
		Country      form.StringField[CompanyBillingRequest]
		VatNumber    form.StringField[CompanyBillingRequest]
		BillingEmail form.StringField[CompanyBillingRequest]
	}{
		AccountType: form.Str[CompanyBillingRequest]("account_type", func(r *CompanyBillingRequest) string {
			return r.AccountType
		}),
		CompanyName: form.Str[CompanyBillingRequest]("company_name", func(r *CompanyBillingRequest) string {
			return r.CompanyName
		}),
		Country: form.Str[CompanyBillingRequest]("country", func(r *CompanyBillingRequest) string {
			return r.Country
		}),
		VatNumber: form.Str[CompanyBillingRequest]("vat_number", func(r *CompanyBillingRequest) string {
			return r.VatNumber
		}),
		BillingEmail: form.Str[CompanyBillingRequest]("billing_email", func(r *CompanyBillingRequest) string {
			return r.BillingEmail
		}),
	}

	AccountTypeSchema := form.Schema[CompanyBillingRequest]{
		rules.Required(CompanyBillingForm.AccountType),
		rules.OneOf[CompanyBillingRequest, string](
			CompanyBillingForm.AccountType.Field,
			"personal",
			"company",
		),
	}

	CompanySchema := form.Schema[CompanyBillingRequest]{
		// company_name is required only for company accounts.
		conditional.RequiredIfStrWithCode(
			CompanyBillingForm.CompanyName,
			func(r *CompanyBillingRequest) bool {
				return r.AccountType == "company"
			},
			CodeCompanyNameRequired,
		),
	}

	VatSchema := form.Schema[CompanyBillingRequest]{
		// vat_number is required only for Slovak company billing.
		conditional.RequiredIfStrWithCode(
			CompanyBillingForm.VatNumber,
			func(r *CompanyBillingRequest) bool {
				return r.AccountType == "company" && r.Country == "SK"
			},
			CodeVatNumberRequired,
		),
	}

	BillingEmailSchema := form.Schema[CompanyBillingRequest]{
		rules.RequiredWithCode(CompanyBillingForm.BillingEmail, CodeBillingEmailRequired),
		rules.EmailWithCode(CompanyBillingForm.BillingEmail, CodeBillingEmailInvalid),
	}

	return form.Rules(
		AccountTypeSchema,
		CompanySchema,
		VatSchema,
		BillingEmailSchema,
	)
}
