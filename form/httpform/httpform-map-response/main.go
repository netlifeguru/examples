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

	r.HandleFunc("/map-response", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in MapResponseRequest

		if !httpform.BindAndValidate(w, req, &in, MapResponseSchema(), 1<<20) {
			fmt.Println("map response validation failed")
			return
		}

		fmt.Println("map response request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "map response validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"email":    "john@example.com",
		"password": "secret123",
	}

	invalidPayload := map[string]any{
		"email":    "invalid-email", // Email fails
		"password": "short",         // MinLen(8) fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/map-response", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/map-response", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
