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

	r.HandleFunc("/required-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in RequiredRulesRequest

		if !httpform.BindAndValidate(w, req, &in, RequiredRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("required rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "required rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"email":       "john@example.com",           // Required
		"name":        "John",                       // RequiredWithCode
		"age":         18,                           // RequiredInt
		"level":       1,                            // RequiredIntWithCode
		"accepted":    true,                         // RequiredBool
		"confirmed":   true,                         // RequiredBoolWithCode
		"birth_date":  "1990-01-01T00:00:00Z",       // RequiredTime
		"started_at":  "2026-01-01T10:00:00Z",       // RequiredTimeWithCode
		"permissions": []string{"read", "write"},    // RequiredSliceStr
		"roles":       []string{"admin", "manager"}, // RequiredSliceStrWithCode
	}

	invalidPayload := map[string]any{
		"email":       "",                     // Required fails
		"name":        "",                     // RequiredWithCode fails
		"age":         0,                      // RequiredInt fails
		"level":       0,                      // RequiredIntWithCode fails
		"accepted":    false,                  // RequiredBool fails
		"confirmed":   false,                  // RequiredBoolWithCode fails
		"birth_date":  "0001-01-01T00:00:00Z", // RequiredTime fails
		"started_at":  "0001-01-01T00:00:00Z", // RequiredTimeWithCode fails
		"permissions": []string{},             // RequiredSliceStr fails
		"roles":       []string{},             // RequiredSliceStrWithCode fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/required-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/required-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
