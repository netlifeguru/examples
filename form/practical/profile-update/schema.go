package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/optional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeNicknameMinLen = form.Code("nickname_min_len")
	CodeBioMaxLen      = form.Code("bio_max_len")
	CodeWebsiteURL     = form.Code("website_url")
	CodeAvatarURL      = form.Code("avatar_url")
)

type ProfileUpdateRequest struct {
	Nickname  string `json:"nickname"`
	Bio       string `json:"bio"`
	Website   string `json:"website"`
	AvatarURL string `json:"avatar_url"`
}

func ProfileUpdateSchema() form.Schema[ProfileUpdateRequest] {
	ProfileUpdateForm := struct {
		Nickname  form.StringField[ProfileUpdateRequest]
		Bio       form.StringField[ProfileUpdateRequest]
		Website   form.StringField[ProfileUpdateRequest]
		AvatarURL form.StringField[ProfileUpdateRequest]
	}{
		Nickname: form.Str[ProfileUpdateRequest]("nickname", func(r *ProfileUpdateRequest) string {
			return r.Nickname
		}),
		Bio: form.Str[ProfileUpdateRequest]("bio", func(r *ProfileUpdateRequest) string {
			return r.Bio
		}),
		Website: form.Str[ProfileUpdateRequest]("website", func(r *ProfileUpdateRequest) string {
			return r.Website
		}),
		AvatarURL: form.Str[ProfileUpdateRequest]("avatar_url", func(r *ProfileUpdateRequest) string {
			return r.AvatarURL
		}),
	}

	NicknameSchema := form.Schema[ProfileUpdateRequest]{
		optional.OptionalString(
			ProfileUpdateForm.Nickname,
			rules.MinLenWithCode(ProfileUpdateForm.Nickname, 3, CodeNicknameMinLen),
			rules.MaxLen(ProfileUpdateForm.Nickname, 30),
		),
	}

	BioSchema := form.Schema[ProfileUpdateRequest]{
		optional.OptionalString(
			ProfileUpdateForm.Bio,
			rules.MaxLenWithCode(ProfileUpdateForm.Bio, 160, CodeBioMaxLen),
		),
	}

	WebsiteSchema := form.Schema[ProfileUpdateRequest]{
		optional.OptionalString(
			ProfileUpdateForm.Website,
			rules.URLWithCode(ProfileUpdateForm.Website, CodeWebsiteURL),
		),
	}

	AvatarSchema := form.Schema[ProfileUpdateRequest]{
		optional.OptionalString(
			ProfileUpdateForm.AvatarURL,
			rules.URLWithCode(ProfileUpdateForm.AvatarURL, CodeAvatarURL),
		),
	}

	return form.Rules(
		NicknameSchema,
		BioSchema,
		WebsiteSchema,
		AvatarSchema,
	)
}
