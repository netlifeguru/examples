package main

import (
	"regexp"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/optional"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeMinTags          = form.Code("min_tags")
	CodeMaxPermissions   = form.Code("max_permissions")
	CodeDistinctFeatures = form.Code("distinct_features")
	CodeItemRequired     = form.Code("item_required")
	CodeItemMinLen       = form.Code("item_min_len")
	CodeItemRegex        = form.Code("item_regex")
)

type OptionalSliceRequest struct {
	Tags        []string `json:"tags"`
	Permissions []string `json:"permissions"`
	Features    []string `json:"features"`
	Codes       []string `json:"codes"`

	RequiredItems []string `json:"required_items"`
	MinLenItems   []string `json:"min_len_items"`
	RegexItems    []string `json:"regex_items"`
	CustomItems   []string `json:"custom_items"`
}

func OptionalSliceSchema() form.Schema[OptionalSliceRequest] {
	OptionalSliceForm := struct {
		Tags        form.SliceStringField[OptionalSliceRequest]
		Permissions form.SliceStringField[OptionalSliceRequest]
		Features    form.SliceStringField[OptionalSliceRequest]
		Codes       form.SliceStringField[OptionalSliceRequest]

		RequiredItems form.SliceStringField[OptionalSliceRequest]
		MinLenItems   form.SliceStringField[OptionalSliceRequest]
		RegexItems    form.SliceStringField[OptionalSliceRequest]
		CustomItems   form.SliceStringField[OptionalSliceRequest]
	}{
		Tags: form.SliceStr[OptionalSliceRequest]("tags", func(r *OptionalSliceRequest) []string {
			return r.Tags
		}),
		Permissions: form.SliceStr[OptionalSliceRequest]("permissions", func(r *OptionalSliceRequest) []string {
			return r.Permissions
		}),
		Features: form.SliceStr[OptionalSliceRequest]("features", func(r *OptionalSliceRequest) []string {
			return r.Features
		}),
		Codes: form.SliceStr[OptionalSliceRequest]("codes", func(r *OptionalSliceRequest) []string {
			return r.Codes
		}),
		RequiredItems: form.SliceStr[OptionalSliceRequest]("required_items", func(r *OptionalSliceRequest) []string {
			return r.RequiredItems
		}),
		MinLenItems: form.SliceStr[OptionalSliceRequest]("min_len_items", func(r *OptionalSliceRequest) []string {
			return r.MinLenItems
		}),
		RegexItems: form.SliceStr[OptionalSliceRequest]("regex_items", func(r *OptionalSliceRequest) []string {
			return r.RegexItems
		}),
		CustomItems: form.SliceStr[OptionalSliceRequest]("custom_items", func(r *OptionalSliceRequest) []string {
			return r.CustomItems
		}),
	}

	codeRegex := regexp.MustCompile(`^[A-Z]{3}-[0-9]{3}$`)

	return form.Schema[OptionalSliceRequest]{
		// OptionalSlice skips nested validation when the slice is empty.

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.Tags.Field,
			rules.MinItemsStrWithCode(OptionalSliceForm.Tags, 2, CodeMinTags),
		),

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.Permissions.Field,
			rules.MaxItemsStrWithCode(OptionalSliceForm.Permissions, 3, CodeMaxPermissions),
		),

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.Features.Field,
			rules.DistinctStrWithCode(OptionalSliceForm.Features, CodeDistinctFeatures),
		),

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.Codes.Field,
			rules.MinItemsStr(OptionalSliceForm.Codes, 2),
			rules.MaxItemsStr(OptionalSliceForm.Codes, 3),
			rules.DistinctStr(OptionalSliceForm.Codes),
		),

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.RequiredItems.Field,
			rules.EachStr(
				OptionalSliceForm.RequiredItems,
				rules.VRequiredWithCode(CodeItemRequired),
			),
		),

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.MinLenItems.Field,
			rules.EachStr(
				OptionalSliceForm.MinLenItems,
				rules.VMinLenWithCode(3, CodeItemMinLen),
			),
		),

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.RegexItems.Field,
			rules.EachStr(
				OptionalSliceForm.RegexItems,
				rules.VRegexWithCode(codeRegex, CodeItemRegex),
			),
		),

		optional.OptionalSlice[OptionalSliceRequest, string](
			OptionalSliceForm.CustomItems.Field,
			rules.EachStr(
				OptionalSliceForm.CustomItems,
				rules.VRequiredWithCode(CodeItemRequired),
				rules.VMinLenWithCode(3, CodeItemMinLen),
				rules.VRegexWithCode(codeRegex, CodeItemRegex),
			),
		),
	}
}
