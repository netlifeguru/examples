package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeMinAge       = form.Code("min_age")
	CodeMaxQuantity  = form.Code("max_quantity")
	CodeScoreBetween = form.Code("score_between")
)

type IntRulesRequest struct {
	Age      int `json:"age"`
	Level    int `json:"level"`
	Quantity int `json:"quantity"`
	Discount int `json:"discount"`
	Rating   int `json:"rating"`
	Score    int `json:"score"`
}

func IntRulesSchema() form.Schema[IntRulesRequest] {
	IntForm := struct {
		Age      form.IntField[IntRulesRequest]
		Level    form.IntField[IntRulesRequest]
		Quantity form.IntField[IntRulesRequest]
		Discount form.IntField[IntRulesRequest]
		Rating   form.IntField[IntRulesRequest]
		Score    form.IntField[IntRulesRequest]
	}{
		Age: form.Int[IntRulesRequest]("age", func(r *IntRulesRequest) int {
			return r.Age
		}),
		Level: form.Int[IntRulesRequest]("level", func(r *IntRulesRequest) int {
			return r.Level
		}),
		Quantity: form.Int[IntRulesRequest]("quantity", func(r *IntRulesRequest) int {
			return r.Quantity
		}),
		Discount: form.Int[IntRulesRequest]("discount", func(r *IntRulesRequest) int {
			return r.Discount
		}),
		Rating: form.Int[IntRulesRequest]("rating", func(r *IntRulesRequest) int {
			return r.Rating
		}),
		Score: form.Int[IntRulesRequest]("score", func(r *IntRulesRequest) int {
			return r.Score
		}),
	}

	return form.Schema[IntRulesRequest]{
		// Min validates that the value is greater than or equal to the minimum.
		rules.Min(IntForm.Age, 18),

		// MinWithCode validates minimum value and returns a custom error code.
		rules.MinWithCode(IntForm.Level, 1, CodeMinAge),

		// Max validates that the value is less than or equal to the maximum.
		rules.Max(IntForm.Quantity, 100),

		// MaxWithCode validates maximum value and returns a custom error code.
		rules.MaxWithCode(IntForm.Discount, 50, CodeMaxQuantity),

		// BetweenInt validates that the value is between min and max, inclusive.
		rules.BetweenInt(IntForm.Rating, 1, 5),

		// BetweenIntWithCode validates range and returns a custom error code.
		rules.BetweenIntWithCode(IntForm.Score, 0, 100, CodeScoreBetween),
	}
}
