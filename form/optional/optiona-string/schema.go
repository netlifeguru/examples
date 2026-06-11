package main

import (
	"regexp"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/optional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeNicknameMinLen     = form.Code("nickname_min_len")
	CodeBioMaxLen          = form.Code("bio_max_len")
	CodeSlugRegex          = form.Code("slug_regex")
	CodeDescriptionContent = form.Code("description_content")
	CodeUserIDPrefix       = form.Code("user_id_prefix")
	CodeImageSuffix        = form.Code("image_suffix")
	CodeHandleAlphaDash    = form.Code("handle_alphadash")
	CodeLowercase          = form.Code("lowercase_required")
	CodeUppercase          = form.Code("uppercase_required")
)

type OptionalStringRequest struct {
	Nickname       string `json:"nickname"`
	Bio            string `json:"bio"`
	Code           string `json:"code"`
	Email          string `json:"email"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	UserID         string `json:"user_id"`
	ImageFile      string `json:"image_file"`
	FirstName      string `json:"first_name"`
	Handle         string `json:"handle"`
	LowercaseValue string `json:"lowercase_value"`
	UppercaseValue string `json:"uppercase_value"`
}

func OptionalStringSchema() form.Schema[OptionalStringRequest] {
	OptionalStringForm := struct {
		Nickname       form.StringField[OptionalStringRequest]
		Bio            form.StringField[OptionalStringRequest]
		Code           form.StringField[OptionalStringRequest]
		Email          form.StringField[OptionalStringRequest]
		Slug           form.StringField[OptionalStringRequest]
		Description    form.StringField[OptionalStringRequest]
		UserID         form.StringField[OptionalStringRequest]
		ImageFile      form.StringField[OptionalStringRequest]
		FirstName      form.StringField[OptionalStringRequest]
		Handle         form.StringField[OptionalStringRequest]
		LowercaseValue form.StringField[OptionalStringRequest]
		UppercaseValue form.StringField[OptionalStringRequest]
	}{
		Nickname: form.Str[OptionalStringRequest]("nickname", func(r *OptionalStringRequest) string {
			return r.Nickname
		}),
		Bio: form.Str[OptionalStringRequest]("bio", func(r *OptionalStringRequest) string {
			return r.Bio
		}),
		Code: form.Str[OptionalStringRequest]("code", func(r *OptionalStringRequest) string {
			return r.Code
		}),
		Email: form.Str[OptionalStringRequest]("email", func(r *OptionalStringRequest) string {
			return r.Email
		}),
		Slug: form.Str[OptionalStringRequest]("slug", func(r *OptionalStringRequest) string {
			return r.Slug
		}),
		Description: form.Str[OptionalStringRequest]("description", func(r *OptionalStringRequest) string {
			return r.Description
		}),
		UserID: form.Str[OptionalStringRequest]("user_id", func(r *OptionalStringRequest) string {
			return r.UserID
		}),
		ImageFile: form.Str[OptionalStringRequest]("image_file", func(r *OptionalStringRequest) string {
			return r.ImageFile
		}),
		FirstName: form.Str[OptionalStringRequest]("first_name", func(r *OptionalStringRequest) string {
			return r.FirstName
		}),
		Handle: form.Str[OptionalStringRequest]("handle", func(r *OptionalStringRequest) string {
			return r.Handle
		}),
		LowercaseValue: form.Str[OptionalStringRequest]("lowercase_value", func(r *OptionalStringRequest) string {
			return r.LowercaseValue
		}),
		UppercaseValue: form.Str[OptionalStringRequest]("uppercase_value", func(r *OptionalStringRequest) string {
			return r.UppercaseValue
		}),
	}

	slugRegex := regexp.MustCompile(`^[a-z0-9-]+$`)

	return form.Schema[OptionalStringRequest]{
		// OptionalString skips nested validation when the string is empty.

		optional.OptionalString(
			OptionalStringForm.Nickname,
			rules.MinLenWithCode(OptionalStringForm.Nickname, 3, CodeNicknameMinLen),
		),

		optional.OptionalString(
			OptionalStringForm.Bio,
			rules.MaxLenWithCode(OptionalStringForm.Bio, 20, CodeBioMaxLen),
		),

		optional.OptionalString(
			OptionalStringForm.Code,
			rules.Len(OptionalStringForm.Code, 6),
		),

		optional.OptionalString(
			OptionalStringForm.Email,
			rules.Email(OptionalStringForm.Email),
		),

		optional.OptionalString(
			OptionalStringForm.Slug,
			rules.RegexWithCode(OptionalStringForm.Slug, slugRegex, CodeSlugRegex),
		),

		optional.OptionalString(
			OptionalStringForm.Description,
			rules.ContainsWithCode(OptionalStringForm.Description, "form", CodeDescriptionContent),
		),

		optional.OptionalString(
			OptionalStringForm.UserID,
			rules.StartsWithWithCode(OptionalStringForm.UserID, "usr_", CodeUserIDPrefix),
		),

		optional.OptionalString(
			OptionalStringForm.ImageFile,
			rules.EndsWithWithCode(OptionalStringForm.ImageFile, ".png", CodeImageSuffix),
		),

		optional.OptionalString(
			OptionalStringForm.FirstName,
			rules.Alpha(OptionalStringForm.FirstName),
		),

		optional.OptionalString(
			OptionalStringForm.Handle,
			rules.AlphaDashWithCode(OptionalStringForm.Handle, CodeHandleAlphaDash),
		),

		optional.OptionalString(
			OptionalStringForm.LowercaseValue,
			rules.LowercaseWithCode(OptionalStringForm.LowercaseValue, CodeLowercase),
		),

		optional.OptionalString(
			OptionalStringForm.UppercaseValue,
			rules.UppercaseWithCode(OptionalStringForm.UppercaseValue, CodeUppercase),
		),
	}
}
