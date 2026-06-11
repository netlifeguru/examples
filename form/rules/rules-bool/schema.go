package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeMustBeAccepted = form.Code("must_be_accepted")
	CodeMustBeDisabled = form.Code("must_be_disabled")
	CodeMustBePublic   = form.Code("must_be_public")
)

type BoolRulesRequest struct {
	TermsAccepted        bool `json:"terms_accepted"`
	MarketingOptOut      bool `json:"marketing_opt_out"`
	Admin                bool `json:"admin"`
	PublicProfile        bool `json:"public_profile"`
	NewsletterSubscribed bool `json:"newsletter_subscribed"`
	FeatureEnabled       bool `json:"feature_enabled"`
	BooleanValue         bool `json:"boolean_value"`
}

func BoolRulesSchema() form.Schema[BoolRulesRequest] {
	BoolForm := struct {
		TermsAccepted        form.BoolField[BoolRulesRequest]
		MarketingOptOut      form.BoolField[BoolRulesRequest]
		Admin                form.BoolField[BoolRulesRequest]
		PublicProfile        form.BoolField[BoolRulesRequest]
		NewsletterSubscribed form.BoolField[BoolRulesRequest]
		FeatureEnabled       form.BoolField[BoolRulesRequest]
		BooleanValue         form.BoolField[BoolRulesRequest]
	}{
		TermsAccepted: form.Bool[BoolRulesRequest]("terms_accepted", func(r *BoolRulesRequest) bool {
			return r.TermsAccepted
		}),
		MarketingOptOut: form.Bool[BoolRulesRequest]("marketing_opt_out", func(r *BoolRulesRequest) bool {
			return r.MarketingOptOut
		}),
		Admin: form.Bool[BoolRulesRequest]("admin", func(r *BoolRulesRequest) bool {
			return r.Admin
		}),
		PublicProfile: form.Bool[BoolRulesRequest]("public_profile", func(r *BoolRulesRequest) bool {
			return r.PublicProfile
		}),
		NewsletterSubscribed: form.Bool[BoolRulesRequest]("newsletter_subscribed", func(r *BoolRulesRequest) bool {
			return r.NewsletterSubscribed
		}),
		FeatureEnabled: form.Bool[BoolRulesRequest]("feature_enabled", func(r *BoolRulesRequest) bool {
			return r.FeatureEnabled
		}),
		BooleanValue: form.Bool[BoolRulesRequest]("boolean_value", func(r *BoolRulesRequest) bool {
			return r.BooleanValue
		}),
	}

	return form.Schema[BoolRulesRequest]{
		rules.IsTrue(BoolForm.TermsAccepted),
		rules.IsTrueWithCode(BoolForm.MarketingOptOut, CodeMustBeAccepted),

		rules.IsFalse(BoolForm.Admin),
		rules.IsFalseWithCode(BoolForm.PublicProfile, CodeMustBeDisabled),

		rules.BoolEquals(BoolForm.NewsletterSubscribed, true),
		rules.BoolEqualsWithCode(BoolForm.FeatureEnabled, false, CodeMustBePublic),

		// IsBool is a no-op for Go typed inputs.
		// Type errors are handled by JSON decoding before validation.
		rules.IsBool(BoolForm.BooleanValue),
	}
}
