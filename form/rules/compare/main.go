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

	r.HandleFunc("/compare-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in CompareRequest

		if !httpform.BindAndValidate(w, req, &in, CompareSchema(), 1<<20) {
			fmt.Println("compare validation failed")
			return
		}

		fmt.Println("compare validation passed:", in)

		w.Header().Set("Content-Type", "application/json")

		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "compare validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"password":         "secret123",
		"confirm_password": "secret123",
		"min_price":        10,
		"max_price":        100,
	}

	invalidPayload := map[string]any{
		"password":         "secret123",
		"confirm_password": "secret321",
		"min_price":        100,
		"max_price":        10,
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/compare-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/compare-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
