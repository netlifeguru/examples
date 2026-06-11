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

	r.HandleFunc("/time-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in TimeRulesRequest

		if !httpform.BindAndValidate(w, req, &in, TimeRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("time rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "time rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"start_at":         "2026-01-02T00:00:00Z", // After(2026-01-01)
		"available_from":   "2026-01-02T00:00:00Z", // AfterWithCode(2026-01-01)
		"deadline":         "2026-12-30T23:59:59Z", // Before(2026-12-31 23:59:59)
		"expires_at":       "2026-12-30T23:59:59Z", // BeforeWithCode(2026-12-31 23:59:59)
		"event_date":       "2026-06-15T12:00:00Z", // BetweenTime(2026-06-01, 2026-06-30)
		"reservation_date": "2026-06-20T12:00:00Z", // BetweenTimeWithCode(2026-06-01, 2026-06-30)
	}

	invalidPayload := map[string]any{
		"start_at":         "2026-01-01T00:00:00Z", // After fails: must be strictly after
		"available_from":   "2025-12-31T23:59:59Z", // AfterWithCode fails
		"deadline":         "2027-01-01T00:00:00Z", // Before fails
		"expires_at":       "2026-12-31T23:59:59Z", // BeforeWithCode fails: must be strictly before
		"event_date":       "2026-07-01T00:00:00Z", // BetweenTime fails
		"reservation_date": "2026-05-31T23:59:59Z", // BetweenTimeWithCode fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/time-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/time-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
