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

	r.HandleFunc("/optional-int", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in OptionalIntRequest

		if !httpform.BindAndValidate(w, req, &in, OptionalIntSchema(), 1<<20) {
			fmt.Println("optional int validation failed")
			return
		}

		fmt.Println("optional int request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "optional int validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"age":      18,  // MinOpt(18)
		"score":    50,  // MinOptWithCode(50)
		"attempts": 3,   // MaxOpt(3)
		"limit":    100, // MaxOptWithCode(100)
		"rating":   5,   // BetweenIntOpt(1, 5)
		"priority": 10,  // BetweenIntOptWithCode(1, 10)
	}

	invalidPayload := map[string]any{
		"age":      17,  // MinOpt fails
		"score":    49,  // MinOptWithCode fails
		"attempts": 4,   // MaxOpt fails
		"limit":    101, // MaxOptWithCode fails
		"rating":   6,   // BetweenIntOpt fails
		"priority": 11,  // BetweenIntOptWithCode fails
	}

	skippedPayload := map[string]any{
		"age":      nil, // skipped
		"score":    nil, // skipped
		"attempts": nil, // skipped
		"limit":    nil, // skipped
		"rating":   nil, // skipped
		"priority": nil, // skipped
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/optional-int", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/optional-int", invalidPayload)

	fmt.Println("\n--- Skipped optional validation request ---")
	form.SendTestPost(":8080/optional-int", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
