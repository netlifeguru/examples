package main

import (
	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/optional"
)

const (
	CodeMinPrice      = form.Code("min_price")
	CodeMaxDiscount   = form.Code("max_discount")
	CodeRatingBetween = form.Code("rating_between")
)

type OptionalFloat64Request struct {
	Price       *float64 `json:"price"`
	MinPrice    *float64 `json:"min_price"`
	Discount    *float64 `json:"discount"`
	MaxDiscount *float64 `json:"max_discount"`
	Rating      *float64 `json:"rating"`
	Score       *float64 `json:"score"`
}

func OptionalFloat64Schema() form.Schema[OptionalFloat64Request] {
	OptionalFloat64Form := struct {
		Price       form.OptFloat64Field[OptionalFloat64Request]
		MinPrice    form.OptFloat64Field[OptionalFloat64Request]
		Discount    form.OptFloat64Field[OptionalFloat64Request]
		MaxDiscount form.OptFloat64Field[OptionalFloat64Request]
		Rating      form.OptFloat64Field[OptionalFloat64Request]
		Score       form.OptFloat64Field[OptionalFloat64Request]
	}{
		Price: form.OptFloat64[OptionalFloat64Request]("price", func(r *OptionalFloat64Request) *float64 {
			return r.Price
		}),
		MinPrice: form.OptFloat64[OptionalFloat64Request]("min_price", func(r *OptionalFloat64Request) *float64 {
			return r.MinPrice
		}),
		Discount: form.OptFloat64[OptionalFloat64Request]("discount", func(r *OptionalFloat64Request) *float64 {
			return r.Discount
		}),
		MaxDiscount: form.OptFloat64[OptionalFloat64Request]("max_discount", func(r *OptionalFloat64Request) *float64 {
			return r.MaxDiscount
		}),
		Rating: form.OptFloat64[OptionalFloat64Request]("rating", func(r *OptionalFloat64Request) *float64 {
			return r.Rating
		}),
		Score: form.OptFloat64[OptionalFloat64Request]("score", func(r *OptionalFloat64Request) *float64 {
			return r.Score
		}),
	}

	return form.Schema[OptionalFloat64Request]{
		// Optional float64 rules skip validation when the pointer is nil.

		optional.MinFloat64Opt(OptionalFloat64Form.Price, 0.01),
		optional.MinFloat64OptWithCode(OptionalFloat64Form.MinPrice, 10.50, CodeMinPrice),

		optional.MaxFloat64Opt(OptionalFloat64Form.Discount, 50.00),
		optional.MaxFloat64OptWithCode(OptionalFloat64Form.MaxDiscount, 75.50, CodeMaxDiscount),

		optional.BetweenFloat64Opt(OptionalFloat64Form.Rating, 1.0, 5.0),
		optional.BetweenFloat64OptWithCode(OptionalFloat64Form.Score, 0.0, 100.0, CodeRatingBetween),
	}
}
