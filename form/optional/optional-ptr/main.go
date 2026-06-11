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

	r.HandleFunc("/optional-ptr", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in OptionalPtrRequest

		if !httpform.BindAndValidate(w, req, &in, OptionalPtrSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("optional ptr request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "optional ptr validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"age":              18,                     // MinOpt(18)
		"score":            50,                     // MinOptWithCode(50)
		"attempts":         3,                      // MaxOpt(3)
		"limit":            100,                    // MaxOptWithCode(100)
		"rating":           5,                      // BetweenIntOpt(1, 5)
		"priority":         10,                     // BetweenIntOptWithCode(1, 10)
		"start_at":         "2026-01-02T00:00:00Z", // AfterOpt
		"available_from":   "2026-01-02T00:00:00Z", // AfterOptWithCode
		"deadline":         "2026-12-30T23:59:59Z", // BeforeOpt
		"expires_at":       "2026-12-30T23:59:59Z", // BeforeOptWithCode
		"event_date":       "2026-06-15T12:00:00Z", // BetweenTimeOpt
		"reservation_date": "2026-06-20T12:00:00Z", // BetweenTimeOptWithCode
		"nickname":         "john",                 // OptionalString + MinLen
		"tags":             []string{"go", "form"}, // OptionalSlice + MinItemsStr + DistinctStr
		"profile_bio":      "I use form package",   // OptionalPtr triggers profile_title requirement
		"profile_title":    "Developer",
	}

	invalidPayload := map[string]any{
		"age":              17,                     // MinOpt fails
		"score":            49,                     // MinOptWithCode fails
		"attempts":         4,                      // MaxOpt fails
		"limit":            101,                    // MaxOptWithCode fails
		"rating":           6,                      // BetweenIntOpt fails
		"priority":         11,                     // BetweenIntOptWithCode fails
		"start_at":         "2026-01-01T00:00:00Z", // AfterOpt fails: must be strictly after
		"available_from":   "2025-12-31T23:59:59Z", // AfterOptWithCode fails
		"deadline":         "2027-01-01T00:00:00Z", // BeforeOpt fails
		"expires_at":       "2026-12-31T23:59:59Z", // BeforeOptWithCode fails: must be strictly before
		"event_date":       "2026-07-01T00:00:00Z", // BetweenTimeOpt fails
		"reservation_date": "2026-05-31T23:59:59Z", // BetweenTimeOptWithCode fails
		"nickname":         "jo",                   // OptionalString triggers MinLen and fails
		"tags":             []string{"go", "go"},   // OptionalSlice triggers DistinctStr and fails
		"profile_bio":      "I use form package",   // OptionalPtr triggers profile_title requirement
		"profile_title":    "",
	}

	skippedPayload := map[string]any{
		"age":              nil,        // skipped
		"score":            nil,        // skipped
		"attempts":         nil,        // skipped
		"limit":            nil,        // skipped
		"rating":           nil,        // skipped
		"priority":         nil,        // skipped
		"start_at":         nil,        // skipped
		"available_from":   nil,        // skipped
		"deadline":         nil,        // skipped
		"expires_at":       nil,        // skipped
		"event_date":       nil,        // skipped
		"reservation_date": nil,        // skipped
		"nickname":         "",         // OptionalString skipped
		"tags":             []string{}, // OptionalSlice skipped
		"profile_bio":      nil,        // OptionalPtr skipped
		"profile_title":    "",
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/optional-ptr", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/optional-ptr", invalidPayload)

	fmt.Println("\n--- Skipped optional validation request ---")
	form.SendTestPost(":8080/optional-ptr", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
