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

	r.HandleFunc("/schema-composition", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in SchemaCompositionRequest

		if !httpform.BindAndValidate(w, req, &in, SchemaCompositionSchema(), 1<<20) {
			fmt.Println("schema composition validation failed")
			return
		}

		fmt.Println("schema composition request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "schema composition validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"name":     "John",
		"email":    "john@example.com",
		"password": "secret123",
		"age":      18,
	}

	invalidPayload := map[string]any{
		"name":     "J",             // MinLen(2) fails
		"email":    "invalid-email", // EmailWithCode fails
		"password": "short",         // MinLenWithCode(8) fails
		"age":      17,              // MinWithCode(18) fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/schema-composition", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/schema-composition", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
