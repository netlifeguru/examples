package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

type PostRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PostSchema() form.Schema[PostRequest] {
	var PostForm = struct {
		Name form.StringField[PostRequest]
		Age  form.IntField[PostRequest]
	}{
		Name: form.Str[PostRequest]("name", func(r *PostRequest) string { return r.Name }),
		Age:  form.Int[PostRequest]("age", func(r *PostRequest) int { return r.Age }),
	}

	var NameSchema = form.Schema[PostRequest]{
		rules.Required(PostForm.Name),
		rules.MinLen(PostForm.Name, 5),
	}

	var AgeSchema = form.Schema[PostRequest]{
		rules.RequiredInt(PostForm.Age),
		rules.Min(PostForm.Age, 8),
	}

	return form.Rules(
		NameSchema,
		AgeSchema,
	)
}
