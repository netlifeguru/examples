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

	r.HandleFunc("/int-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in IntRulesRequest

		if !httpform.BindAndValidate(w, req, &in, IntRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("int rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "int rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"age":      18,  // Min(18)
		"level":    1,   // MinWithCode(1)
		"quantity": 100, // Max(100)
		"discount": 50,  // MaxWithCode(50)
		"rating":   5,   // BetweenInt(1, 5)
		"score":    100, // BetweenIntWithCode(0, 100)
	}

	invalidPayload := map[string]any{
		"age":      17,  // Min(18) fails
		"level":    0,   // MinWithCode(1) fails
		"quantity": 101, // Max(100) fails
		"discount": 51,  // MaxWithCode(50) fails
		"rating":   6,   // BetweenInt(1, 5) fails
		"score":    101, // BetweenIntWithCode(0, 100) fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/int-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/int-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
