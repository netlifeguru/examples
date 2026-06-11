package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeEmailRequired = form.Code("email_required")
)

type UniqueCodesRequest struct {
	Email string `json:"email"`
}

func UniqueCodesSchema() form.Schema[UniqueCodesRequest] {
	UniqueCodesForm := struct {
		Email form.StringField[UniqueCodesRequest]
	}{
		Email: form.Str[UniqueCodesRequest]("email", func(r *UniqueCodesRequest) string {
			return r.Email
		}),
	}

	return form.Schema[UniqueCodesRequest]{
		// Same rule intentionally added twice to demonstrate duplicated error codes.
		rules.RequiredWithCode(UniqueCodesForm.Email, CodeEmailRequired),
		rules.RequiredWithCode(UniqueCodesForm.Email, CodeEmailRequired),
	}
}
