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

	r.HandleFunc("/optional-time", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in OptionalTimeRequest

		if !httpform.BindAndValidate(w, req, &in, OptionalTimeSchema(), 1<<20) {
			fmt.Println("optional time validation failed")
			return
		}

		fmt.Println("optional time request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "optional time validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"start_at":         "2026-01-02T00:00:00Z", // AfterOpt
		"available_from":   "2026-01-02T00:00:00Z", // AfterOptWithCode
		"deadline":         "2026-12-30T23:59:59Z", // BeforeOpt
		"expires_at":       "2026-12-30T23:59:59Z", // BeforeOptWithCode
		"event_date":       "2026-06-15T12:00:00Z", // BetweenTimeOpt
		"reservation_date": "2026-06-20T12:00:00Z", // BetweenTimeOptWithCode
	}

	invalidPayload := map[string]any{
		"start_at":         "2026-01-01T00:00:00Z", // AfterOpt fails
		"available_from":   "2025-12-31T23:59:59Z", // AfterOptWithCode fails
		"deadline":         "2027-01-01T00:00:00Z", // BeforeOpt fails
		"expires_at":       "2026-12-31T23:59:59Z", // BeforeOptWithCode fails
		"event_date":       "2026-07-01T00:00:00Z", // BetweenTimeOpt fails
		"reservation_date": "2026-05-31T23:59:59Z", // BetweenTimeOptWithCode fails
	}

	skippedPayload := map[string]any{
		"start_at":         nil, // skipped
		"available_from":   nil, // skipped
		"deadline":         nil, // skipped
		"expires_at":       nil, // skipped
		"event_date":       nil, // skipped
		"reservation_date": nil, // skipped
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/optional-time", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/optional-time", invalidPayload)

	fmt.Println("\n--- Skipped optional validation request ---")
	form.SendTestPost(":8080/optional-time", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
