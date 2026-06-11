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

	r.HandleFunc("/bool-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in BoolRulesRequest

		if !httpform.BindAndValidate(w, req, &in, BoolRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("bool rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "bool rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"terms_accepted":        true,  // IsTrue
		"marketing_opt_out":     true,  // IsTrueWithCode
		"admin":                 false, // IsFalse
		"public_profile":        false, // IsFalseWithCode
		"newsletter_subscribed": true,  // BoolEquals(true)
		"feature_enabled":       false, // BoolEqualsWithCode(false)
		"boolean_value":         false, // IsBool is no-op
	}

	invalidPayload := map[string]any{
		"terms_accepted":        false, // IsTrue fails
		"marketing_opt_out":     false, // IsTrueWithCode fails
		"admin":                 true,  // IsFalse fails
		"public_profile":        true,  // IsFalseWithCode fails
		"newsletter_subscribed": false, // BoolEquals(true) fails
		"feature_enabled":       true,  // BoolEqualsWithCode(false) fails
		"boolean_value":         false, // IsBool is no-op
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/bool-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/bool-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
