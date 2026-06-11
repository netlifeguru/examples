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

	r.HandleFunc("/float64-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in Float64RulesRequest

		if !httpform.BindAndValidate(w, req, &in, Float64RulesSchema(), 1<<20) {
			fmt.Println("float64 rules validation failed")
			return
		}

		fmt.Println("float64 rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "float64 rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"price":        0.01,  // MinFloat64(0.01)
		"min_price":    10.50, // MinFloat64WithCode(10.50)
		"discount":     50.00, // MaxFloat64(50.00)
		"max_discount": 75.50, // MaxFloat64WithCode(75.50)
		"rating":       5.0,   // BetweenFloat64(1.0, 5.0)
		"score":        100.0, // BetweenFloat64WithCode(0.0, 100.0)
	}

	invalidPayload := map[string]any{
		"price":        0.00,  // MinFloat64 fails
		"min_price":    10.49, // MinFloat64WithCode fails
		"discount":     50.01, // MaxFloat64 fails
		"max_discount": 75.51, // MaxFloat64WithCode fails
		"rating":       5.01,  // BetweenFloat64 fails
		"score":        -0.01, // BetweenFloat64WithCode fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/float64-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/float64-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
