package main

import (
	"regexp"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeCustomMinLen     = form.Code("custom_min_len")
	CodeCustomMaxLen     = form.Code("custom_max_len")
	CodeCustomLen        = form.Code("custom_len")
	CodeCustomEmail      = form.Code("custom_email")
	CodeCustomRegex      = form.Code("custom_regex")
	CodeCustomNotRegex   = form.Code("custom_not_regex")
	CodeCustomContains   = form.Code("custom_contains")
	CodeCustomStartsWith = form.Code("custom_starts_with")
	CodeCustomEndsWith   = form.Code("custom_ends_with")
	CodeCustomAlpha      = form.Code("custom_alpha")
	CodeCustomAlphaNum   = form.Code("custom_alnum")
	CodeCustomAlphaDash  = form.Code("custom_alphadash")
	CodeCustomLowercase  = form.Code("custom_lowercase")
	CodeCustomUppercase  = form.Code("custom_uppercase")
)

type StringRulesRequest struct {
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Title           string `json:"title"`
	ShortTitle      string `json:"short_title"`
	PIN             string `json:"pin"`
	AccessCode      string `json:"access_code"`
	Email           string `json:"email"`
	ContactEmail    string `json:"contact_email"`
	Slug            string `json:"slug"`
	ProductCode     string `json:"product_code"`
	NoSpaces        string `json:"no_spaces"`
	Path            string `json:"path"`
	Bio             string `json:"bio"`
	Description     string `json:"description"`
	UserID          string `json:"user_id"`
	Token           string `json:"token"`
	ImageFile       string `json:"image_file"`
	DocumentFile    string `json:"document_file"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	AlphaNumCode    string `json:"alphanum_code"`
	Reference       string `json:"reference"`
	Handle          string `json:"handle"`
	ProjectKey      string `json:"project_key"`
	LowercaseValue  string `json:"lowercase_value"`
	LowercaseCustom string `json:"lowercase_custom"`
	UppercaseValue  string `json:"uppercase_value"`
	UppercaseCustom string `json:"uppercase_custom"`
}

func StringRulesSchema() form.Schema[StringRulesRequest] {
	StringForm := struct {
		Username        form.StringField[StringRulesRequest]
		Nickname        form.StringField[StringRulesRequest]
		Title           form.StringField[StringRulesRequest]
		ShortTitle      form.StringField[StringRulesRequest]
		PIN             form.StringField[StringRulesRequest]
		AccessCode      form.StringField[StringRulesRequest]
		Email           form.StringField[StringRulesRequest]
		ContactEmail    form.StringField[StringRulesRequest]
		Slug            form.StringField[StringRulesRequest]
		ProductCode     form.StringField[StringRulesRequest]
		NoSpaces        form.StringField[StringRulesRequest]
		Path            form.StringField[StringRulesRequest]
		Bio             form.StringField[StringRulesRequest]
		Description     form.StringField[StringRulesRequest]
		UserID          form.StringField[StringRulesRequest]
		Token           form.StringField[StringRulesRequest]
		ImageFile       form.StringField[StringRulesRequest]
		DocumentFile    form.StringField[StringRulesRequest]
		FirstName       form.StringField[StringRulesRequest]
		LastName        form.StringField[StringRulesRequest]
		AlphaNumCode    form.StringField[StringRulesRequest]
		Reference       form.StringField[StringRulesRequest]
		Handle          form.StringField[StringRulesRequest]
		ProjectKey      form.StringField[StringRulesRequest]
		LowercaseValue  form.StringField[StringRulesRequest]
		LowercaseCustom form.StringField[StringRulesRequest]
		UppercaseValue  form.StringField[StringRulesRequest]
		UppercaseCustom form.StringField[StringRulesRequest]
	}{
		Username:        form.Str[StringRulesRequest]("username", func(r *StringRulesRequest) string { return r.Username }),
		Nickname:        form.Str[StringRulesRequest]("nickname", func(r *StringRulesRequest) string { return r.Nickname }),
		Title:           form.Str[StringRulesRequest]("title", func(r *StringRulesRequest) string { return r.Title }),
		ShortTitle:      form.Str[StringRulesRequest]("short_title", func(r *StringRulesRequest) string { return r.ShortTitle }),
		PIN:             form.Str[StringRulesRequest]("pin", func(r *StringRulesRequest) string { return r.PIN }),
		AccessCode:      form.Str[StringRulesRequest]("access_code", func(r *StringRulesRequest) string { return r.AccessCode }),
		Email:           form.Str[StringRulesRequest]("email", func(r *StringRulesRequest) string { return r.Email }),
		ContactEmail:    form.Str[StringRulesRequest]("contact_email", func(r *StringRulesRequest) string { return r.ContactEmail }),
		Slug:            form.Str[StringRulesRequest]("slug", func(r *StringRulesRequest) string { return r.Slug }),
		ProductCode:     form.Str[StringRulesRequest]("product_code", func(r *StringRulesRequest) string { return r.ProductCode }),
		NoSpaces:        form.Str[StringRulesRequest]("no_spaces", func(r *StringRulesRequest) string { return r.NoSpaces }),
		Path:            form.Str[StringRulesRequest]("path", func(r *StringRulesRequest) string { return r.Path }),
		Bio:             form.Str[StringRulesRequest]("bio", func(r *StringRulesRequest) string { return r.Bio }),
		Description:     form.Str[StringRulesRequest]("description", func(r *StringRulesRequest) string { return r.Description }),
		UserID:          form.Str[StringRulesRequest]("user_id", func(r *StringRulesRequest) string { return r.UserID }),
		Token:           form.Str[StringRulesRequest]("token", func(r *StringRulesRequest) string { return r.Token }),
		ImageFile:       form.Str[StringRulesRequest]("image_file", func(r *StringRulesRequest) string { return r.ImageFile }),
		DocumentFile:    form.Str[StringRulesRequest]("document_file", func(r *StringRulesRequest) string { return r.DocumentFile }),
		FirstName:       form.Str[StringRulesRequest]("first_name", func(r *StringRulesRequest) string { return r.FirstName }),
		LastName:        form.Str[StringRulesRequest]("last_name", func(r *StringRulesRequest) string { return r.LastName }),
		AlphaNumCode:    form.Str[StringRulesRequest]("alphanum_code", func(r *StringRulesRequest) string { return r.AlphaNumCode }),
		Reference:       form.Str[StringRulesRequest]("reference", func(r *StringRulesRequest) string { return r.Reference }),
		Handle:          form.Str[StringRulesRequest]("handle", func(r *StringRulesRequest) string { return r.Handle }),
		ProjectKey:      form.Str[StringRulesRequest]("project_key", func(r *StringRulesRequest) string { return r.ProjectKey }),
		LowercaseValue:  form.Str[StringRulesRequest]("lowercase_value", func(r *StringRulesRequest) string { return r.LowercaseValue }),
		LowercaseCustom: form.Str[StringRulesRequest]("lowercase_custom", func(r *StringRulesRequest) string { return r.LowercaseCustom }),
		UppercaseValue:  form.Str[StringRulesRequest]("uppercase_value", func(r *StringRulesRequest) string { return r.UppercaseValue }),
		UppercaseCustom: form.Str[StringRulesRequest]("uppercase_custom", func(r *StringRulesRequest) string { return r.UppercaseCustom }),
	}

	slugRegex := regexp.MustCompile(`^[a-z0-9-]+$`)
	productCodeRegex := regexp.MustCompile(`^[A-Z]{3}-[0-9]{3}$`)
	spaceRegex := regexp.MustCompile(`\s`)
	forbiddenAdminRegex := regexp.MustCompile(`admin`)

	return form.Schema[StringRulesRequest]{
		rules.MinLen(StringForm.Username, 5),
		rules.MinLenWithCode(StringForm.Nickname, 5, CodeCustomMinLen),

		rules.MaxLen(StringForm.Title, 10),
		rules.MaxLenWithCode(StringForm.ShortTitle, 10, CodeCustomMaxLen),

		rules.Len(StringForm.PIN, 4),
		rules.LenWithCode(StringForm.AccessCode, 6, CodeCustomLen),

		rules.Email(StringForm.Email),
		rules.EmailWithCode(StringForm.ContactEmail, CodeCustomEmail),

		rules.Regex(StringForm.Slug, slugRegex),
		rules.RegexWithCode(StringForm.ProductCode, productCodeRegex, CodeCustomRegex),

		rules.NotRegex(StringForm.NoSpaces, spaceRegex),
		rules.NotRegexWithCode(StringForm.Path, forbiddenAdminRegex, CodeCustomNotRegex),

		rules.Contains(StringForm.Bio, "form"),
		rules.ContainsWithCode(StringForm.Description, "package", CodeCustomContains),

		rules.StartsWith(StringForm.UserID, "usr_"),
		rules.StartsWithWithCode(StringForm.Token, "tok_", CodeCustomStartsWith),

		rules.EndsWith(StringForm.ImageFile, ".png"),
		rules.EndsWithWithCode(StringForm.DocumentFile, ".pdf", CodeCustomEndsWith),

		rules.Alpha(StringForm.FirstName),
		rules.AlphaWithCode(StringForm.LastName, CodeCustomAlpha),

		rules.AlphaNum(StringForm.AlphaNumCode),
		rules.AlphaNumWithCode(StringForm.Reference, CodeCustomAlphaNum),

		rules.AlphaDash(StringForm.Handle),
		rules.AlphaDashWithCode(StringForm.ProjectKey, CodeCustomAlphaDash),

		rules.Lowercase(StringForm.LowercaseValue),
		rules.LowercaseWithCode(StringForm.LowercaseCustom, CodeCustomLowercase),

		rules.Uppercase(StringForm.UppercaseValue),
		rules.UppercaseWithCode(StringForm.UppercaseCustom, CodeCustomUppercase),
	}
}
