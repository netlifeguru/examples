package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/httpform"
	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	r.HandleFunc("/optional-float64", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in OptionalFloat64Request

		if !httpform.BindAndValidate(w, req, &in, OptionalFloat64Schema(), 1<<20) {
			fmt.Println("optional float64 validation failed")
			return
		}

		fmt.Println("optional float64 request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "optional float64 validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"price":        0.01,  // MinFloat64Opt(0.01)
		"min_price":    10.50, // MinFloat64OptWithCode(10.50)
		"discount":     50.00, // MaxFloat64Opt(50.00)
		"max_discount": 75.50, // MaxFloat64OptWithCode(75.50)
		"rating":       5.0,   // BetweenFloat64Opt(1.0, 5.0)
		"score":        100.0, // BetweenFloat64OptWithCode(0.0, 100.0)
	}

	invalidPayload := map[string]any{
		"price":        0.00,  // MinFloat64Opt fails
		"min_price":    10.49, // MinFloat64OptWithCode fails
		"discount":     50.01, // MaxFloat64Opt fails
		"max_discount": 75.51, // MaxFloat64OptWithCode fails
		"rating":       5.01,  // BetweenFloat64Opt fails
		"score":        -0.01, // BetweenFloat64OptWithCode fails
	}

	skippedPayload := map[string]any{
		"price":        nil, // skipped
		"min_price":    nil, // skipped
		"discount":     nil, // skipped
		"max_discount": nil, // skipped
		"rating":       nil, // skipped
		"score":        nil, // skipped
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/optional-float64", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/optional-float64", invalidPayload)

	fmt.Println("\n--- Skipped optional validation request ---")
	form.SendTestPost(":8080/optional-float64", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
