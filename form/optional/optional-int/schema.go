package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/optional"
)

const (
	CodeMinAge        = form.Code("min_age")
	CodeMaxAttempts   = form.Code("max_attempts")
	CodePriorityRange = form.Code("priority_range")
)

type OptionalIntRequest struct {
	Age      *int `json:"age"`
	Score    *int `json:"score"`
	Attempts *int `json:"attempts"`
	Limit    *int `json:"limit"`
	Rating   *int `json:"rating"`
	Priority *int `json:"priority"`
}

func OptionalIntSchema() form.Schema[OptionalIntRequest] {
	OptionalIntForm := struct {
		Age      form.OptIntField[OptionalIntRequest]
		Score    form.OptIntField[OptionalIntRequest]
		Attempts form.OptIntField[OptionalIntRequest]
		Limit    form.OptIntField[OptionalIntRequest]
		Rating   form.OptIntField[OptionalIntRequest]
		Priority form.OptIntField[OptionalIntRequest]
	}{
		Age: form.OptInt[OptionalIntRequest]("age", func(r *OptionalIntRequest) *int {
			return r.Age
		}),
		Score: form.OptInt[OptionalIntRequest]("score", func(r *OptionalIntRequest) *int {
			return r.Score
		}),
		Attempts: form.OptInt[OptionalIntRequest]("attempts", func(r *OptionalIntRequest) *int {
			return r.Attempts
		}),
		Limit: form.OptInt[OptionalIntRequest]("limit", func(r *OptionalIntRequest) *int {
			return r.Limit
		}),
		Rating: form.OptInt[OptionalIntRequest]("rating", func(r *OptionalIntRequest) *int {
			return r.Rating
		}),
		Priority: form.OptInt[OptionalIntRequest]("priority", func(r *OptionalIntRequest) *int {
			return r.Priority
		}),
	}

	return form.Schema[OptionalIntRequest]{
		// Optional int rules skip validation when the pointer is nil.

		optional.MinOpt(OptionalIntForm.Age, 18),
		optional.MinOptWithCode(OptionalIntForm.Score, 50, CodeMinAge),

		optional.MaxOpt(OptionalIntForm.Attempts, 3),
		optional.MaxOptWithCode(OptionalIntForm.Limit, 100, CodeMaxAttempts),

		optional.BetweenIntOpt(OptionalIntForm.Rating, 1, 5),
		optional.BetweenIntOptWithCode(OptionalIntForm.Priority, 1, 10, CodePriorityRange),
	}
}
