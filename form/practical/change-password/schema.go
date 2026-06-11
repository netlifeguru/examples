package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/conditional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeCurrentPasswordRequired = form.Code("current_password_required")
	CodeNewPasswordRequired     = form.Code("new_password_required")
	CodeNewPasswordMinLen       = form.Code("new_password_min_len")
	CodeConfirmPasswordRequired = form.Code("confirm_password_required")
)

type ChangePasswordRequest struct {
	ChangePassword  bool   `json:"change_password"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ChangePasswordSchema() form.Schema[ChangePasswordRequest] {
	ChangePasswordForm := struct {
		ChangePassword  form.BoolField[ChangePasswordRequest]
		CurrentPassword form.StringField[ChangePasswordRequest]
		NewPassword     form.StringField[ChangePasswordRequest]
		ConfirmPassword form.StringField[ChangePasswordRequest]
	}{
		ChangePassword: form.Bool[ChangePasswordRequest]("change_password", func(r *ChangePasswordRequest) bool {
			return r.ChangePassword
		}),
		CurrentPassword: form.Str[ChangePasswordRequest]("current_password", func(r *ChangePasswordRequest) string {
			return r.CurrentPassword
		}),
		NewPassword: form.Str[ChangePasswordRequest]("new_password", func(r *ChangePasswordRequest) string {
			return r.NewPassword
		}),
		ConfirmPassword: form.Str[ChangePasswordRequest]("confirm_password", func(r *ChangePasswordRequest) string {
			return r.ConfirmPassword
		}),
	}

	return form.Schema[ChangePasswordRequest]{
		// Password fields are required only when change_password is true.
		conditional.When(
			func(r *ChangePasswordRequest) bool {
				return r.ChangePassword
			},
			rules.RequiredWithCode(ChangePasswordForm.CurrentPassword, CodeCurrentPasswordRequired),
			rules.RequiredWithCode(ChangePasswordForm.NewPassword, CodeNewPasswordRequired),
			rules.MinLenWithCode(ChangePasswordForm.NewPassword, 8, CodeNewPasswordMinLen),
			rules.RequiredWithCode(ChangePasswordForm.ConfirmPassword, CodeConfirmPasswordRequired),
			rules.MinLen(ChangePasswordForm.ConfirmPassword, 8),
		),
	}
}
