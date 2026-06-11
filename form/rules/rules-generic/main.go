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

	r.HandleFunc("/generic-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in GenericRulesRequest

		if !httpform.BindAndValidate(w, req, &in, GenericRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("generic rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "generic rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"role":     "admin",      // OneOf("admin", "editor", "viewer")
		"status":   "published",  // OneOfWithCode("draft", "published", "archived")
		"plan":     "enterprise", // OneOfWithCode("free", "pro", "enterprise")
		"level":    2,            // OneOf(1, 2, 3)
		"priority": 20,           // OneOfWithCode(10, 20, 30)
		"active":   true,         // OneOf(true)
	}

	invalidPayload := map[string]any{
		"role":     "owner",   // OneOf fails
		"status":   "deleted", // OneOfWithCode fails
		"plan":     "premium", // OneOfWithCode fails
		"level":    9,         // OneOf fails
		"priority": 99,        // OneOfWithCode fails
		"active":   false,     // OneOf(true) fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/generic-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/generic-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
