package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/rules"
)

const (
	CodeMinPrice      = form.Code("min_price")
	CodeMaxDiscount   = form.Code("max_discount")
	CodeRatingBetween = form.Code("rating_between")
)

type Float64RulesRequest struct {
	Price       float64 `json:"price"`
	MinPrice    float64 `json:"min_price"`
	Discount    float64 `json:"discount"`
	MaxDiscount float64 `json:"max_discount"`
	Rating      float64 `json:"rating"`
	Score       float64 `json:"score"`
}

func Float64RulesSchema() form.Schema[Float64RulesRequest] {
	Float64Form := struct {
		Price       form.Float64Field[Float64RulesRequest]
		MinPrice    form.Float64Field[Float64RulesRequest]
		Discount    form.Float64Field[Float64RulesRequest]
		MaxDiscount form.Float64Field[Float64RulesRequest]
		Rating      form.Float64Field[Float64RulesRequest]
		Score       form.Float64Field[Float64RulesRequest]
	}{
		Price: form.Float64[Float64RulesRequest]("price", func(r *Float64RulesRequest) float64 {
			return r.Price
		}),
		MinPrice: form.Float64[Float64RulesRequest]("min_price", func(r *Float64RulesRequest) float64 {
			return r.MinPrice
		}),
		Discount: form.Float64[Float64RulesRequest]("discount", func(r *Float64RulesRequest) float64 {
			return r.Discount
		}),
		MaxDiscount: form.Float64[Float64RulesRequest]("max_discount", func(r *Float64RulesRequest) float64 {
			return r.MaxDiscount
		}),
		Rating: form.Float64[Float64RulesRequest]("rating", func(r *Float64RulesRequest) float64 {
			return r.Rating
		}),
		Score: form.Float64[Float64RulesRequest]("score", func(r *Float64RulesRequest) float64 {
			return r.Score
		}),
	}

	return form.Schema[Float64RulesRequest]{
		// MinFloat64 validates that the value is greater than or equal to the minimum.
		rules.MinFloat64(Float64Form.Price, 0.01),

		// MinFloat64WithCode validates minimum value and returns a custom error code.
		rules.MinFloat64WithCode(Float64Form.MinPrice, 10.50, CodeMinPrice),

		// MaxFloat64 validates that the value is less than or equal to the maximum.
		rules.MaxFloat64(Float64Form.Discount, 50.00),

		// MaxFloat64WithCode validates maximum value and returns a custom error code.
		rules.MaxFloat64WithCode(Float64Form.MaxDiscount, 75.50, CodeMaxDiscount),

		// BetweenFloat64 validates that the value is between min and max, inclusive.
		rules.BetweenFloat64(Float64Form.Rating, 1.0, 5.0),

		// BetweenFloat64WithCode validates range and returns a custom error code.
		rules.BetweenFloat64WithCode(Float64Form.Score, 0.0, 100.0, CodeRatingBetween),
	}
}
