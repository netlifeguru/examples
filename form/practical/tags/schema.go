package main

import (
	"regexp"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeTagsRequired = form.Code("tags_required")
	CodeTagsMin      = form.Code("tags_min")
	CodeTagsMax      = form.Code("tags_max")
	CodeTagsDistinct = form.Code("tags_distinct")
	CodeTagRequired  = form.Code("tag_required")
	CodeTagFormat    = form.Code("tag_format")
)

type TagsRequest struct {
	Tags []string `json:"tags"`
}

func TagsSchema() form.Schema[TagsRequest] {
	TagsForm := struct {
		Tags form.SliceStringField[TagsRequest]
	}{
		Tags: form.SliceStr[TagsRequest]("tags", func(r *TagsRequest) []string {
			return r.Tags
		}),
	}

	tagRegex := regexp.MustCompile(`^[a-z0-9-]+$`)

	TagsSchema := form.Schema[TagsRequest]{
		rules.RequiredSliceStrWithCode(TagsForm.Tags, CodeTagsRequired),
		rules.MinItemsStrWithCode(TagsForm.Tags, 2, CodeTagsMin),
		rules.MaxItemsStrWithCode(TagsForm.Tags, 5, CodeTagsMax),
		rules.DistinctStrWithCode(TagsForm.Tags, CodeTagsDistinct),
		rules.EachStr(
			TagsForm.Tags,
			rules.VRequiredWithCode(CodeTagRequired),
			rules.VRegexWithCode(tagRegex, CodeTagFormat),
		),
	}

	return form.Rules(
		TagsSchema,
	)
}
