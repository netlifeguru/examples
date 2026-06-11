package main

import (
	"regexp"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeRequiredRoles    = form.Code("required_roles")
	CodeMinTags          = form.Code("min_tags")
	CodeMaxPermissions   = form.Code("max_permissions")
	CodeDistinctFeatures = form.Code("distinct_features")
	CodeItemRequired     = form.Code("item_required")
	CodeItemMinLen       = form.Code("item_min_len")
	CodeItemRegex        = form.Code("item_regex")
)

type SliceRulesRequest struct {
	Roles           []string `json:"roles"`
	Permissions     []string `json:"permissions"`
	Tags            []string `json:"tags"`
	Categories      []string `json:"categories"`
	Features        []string `json:"features"`
	UniqueCodes     []string `json:"unique_codes"`
	RequiredItems   []string `json:"required_items"`
	MinLengthItems  []string `json:"min_length_items"`
	RegexItems      []string `json:"regex_items"`
	CustomEachItems []string `json:"custom_each_items"`
}

func SliceRulesSchema() form.Schema[SliceRulesRequest] {
	SliceForm := struct {
		Roles           form.SliceStringField[SliceRulesRequest]
		Permissions     form.SliceStringField[SliceRulesRequest]
		Tags            form.SliceStringField[SliceRulesRequest]
		Categories      form.SliceStringField[SliceRulesRequest]
		Features        form.SliceStringField[SliceRulesRequest]
		UniqueCodes     form.SliceStringField[SliceRulesRequest]
		RequiredItems   form.SliceStringField[SliceRulesRequest]
		MinLengthItems  form.SliceStringField[SliceRulesRequest]
		RegexItems      form.SliceStringField[SliceRulesRequest]
		CustomEachItems form.SliceStringField[SliceRulesRequest]
	}{
		Roles: form.SliceStr[SliceRulesRequest]("roles", func(r *SliceRulesRequest) []string {
			return r.Roles
		}),
		Permissions: form.SliceStr[SliceRulesRequest]("permissions", func(r *SliceRulesRequest) []string {
			return r.Permissions
		}),
		Tags: form.SliceStr[SliceRulesRequest]("tags", func(r *SliceRulesRequest) []string {
			return r.Tags
		}),
		Categories: form.SliceStr[SliceRulesRequest]("categories", func(r *SliceRulesRequest) []string {
			return r.Categories
		}),
		Features: form.SliceStr[SliceRulesRequest]("features", func(r *SliceRulesRequest) []string {
			return r.Features
		}),
		UniqueCodes: form.SliceStr[SliceRulesRequest]("unique_codes", func(r *SliceRulesRequest) []string {
			return r.UniqueCodes
		}),
		RequiredItems: form.SliceStr[SliceRulesRequest]("required_items", func(r *SliceRulesRequest) []string {
			return r.RequiredItems
		}),
		MinLengthItems: form.SliceStr[SliceRulesRequest]("min_length_items", func(r *SliceRulesRequest) []string {
			return r.MinLengthItems
		}),
		RegexItems: form.SliceStr[SliceRulesRequest]("regex_items", func(r *SliceRulesRequest) []string {
			return r.RegexItems
		}),
		CustomEachItems: form.SliceStr[SliceRulesRequest]("custom_each_items", func(r *SliceRulesRequest) []string {
			return r.CustomEachItems
		}),
	}

	codeRegex := regexp.MustCompile(`^[A-Z]{3}-[0-9]{3}$`)

	return form.Schema[SliceRulesRequest]{
		// RequiredSliceStr validates that the slice is not empty.
		rules.RequiredSliceStr(SliceForm.Roles),

		// RequiredSliceStrWithCode validates that the slice is not empty and returns a custom code.
		rules.RequiredSliceStrWithCode(SliceForm.Permissions, CodeRequiredRoles),

		// MinItemsStr validates that the slice has at least n items.
		rules.MinItemsStr(SliceForm.Tags, 2),

		// MinItemsStrWithCode validates minimum item count and returns a custom code.
		rules.MinItemsStrWithCode(SliceForm.Categories, 2, CodeMinTags),

		// MaxItemsStr validates that the slice has at most n items.
		rules.MaxItemsStr(SliceForm.Features, 3),

		// MaxItemsStrWithCode validates maximum item count and returns a custom code.
		rules.MaxItemsStrWithCode(SliceForm.UniqueCodes, 3, CodeMaxPermissions),

		// DistinctStr validates that all values are unique.
		rules.DistinctStr(SliceForm.Features),

		// DistinctStrWithCode validates uniqueness and returns a custom code.
		rules.DistinctStrWithCode(SliceForm.UniqueCodes, CodeDistinctFeatures),

		// EachStr validates every item and reports errors as field[index].
		rules.EachStr(
			SliceForm.RequiredItems,
			rules.VRequired(),
		),

		rules.EachStr(
			SliceForm.MinLengthItems,
			rules.VMinLen(3),
		),

		rules.EachStr(
			SliceForm.RegexItems,
			rules.VRegex(codeRegex),
		),

		rules.EachStr(
			SliceForm.CustomEachItems,
			rules.VRequiredWithCode(CodeItemRequired),
			rules.VMinLenWithCode(3, CodeItemMinLen),
			rules.VRegexWithCode(codeRegex, CodeItemRegex),
		),
	}
}
