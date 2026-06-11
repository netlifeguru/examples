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

	r.HandleFunc("/optional-string", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in OptionalStringRequest

		if !httpform.BindAndValidate(w, req, &in, OptionalStringSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("optional string request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "optional string validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"nickname":        "john",              // OptionalString + MinLenWithCode
		"bio":             "short bio",         // OptionalString + MaxLenWithCode
		"code":            "ABC123",            // OptionalString + Len
		"email":           "john@example.com",  // OptionalString + Email
		"slug":            "hello-world-123",   // OptionalString + RegexWithCode
		"description":     "form package demo", // OptionalString + ContainsWithCode
		"user_id":         "usr_123",           // OptionalString + StartsWithWithCode
		"image_file":      "avatar.png",        // OptionalString + EndsWithWithCode
		"first_name":      "John",              // OptionalString + Alpha
		"handle":          "john_doe-123",      // OptionalString + AlphaDashWithCode
		"lowercase_value": "lowercase",         // OptionalString + LowercaseWithCode
		"uppercase_value": "UPPERCASE",         // OptionalString + UppercaseWithCode
	}

	invalidPayload := map[string]any{
		"nickname":        "jo",                              // MinLenWithCode fails
		"bio":             "this bio is definitely too long", // MaxLenWithCode fails
		"code":            "ABC12",                           // Len fails
		"email":           "invalid-email",                   // Email fails
		"slug":            "Hello World!",                    // RegexWithCode fails
		"description":     "validation demo",                 // ContainsWithCode fails
		"user_id":         "user_123",                        // StartsWithWithCode fails
		"image_file":      "avatar.jpg",                      // EndsWithWithCode fails
		"first_name":      "John123",                         // Alpha fails
		"handle":          "john doe",                        // AlphaDashWithCode fails
		"lowercase_value": "Lowercase",                       // LowercaseWithCode fails
		"uppercase_value": "Uppercase",                       // UppercaseWithCode fails
	}

	skippedPayload := map[string]any{
		"nickname":        "", // skipped
		"bio":             "", // skipped
		"code":            "", // skipped
		"email":           "", // skipped
		"slug":            "", // skipped
		"description":     "", // skipped
		"user_id":         "", // skipped
		"image_file":      "", // skipped
		"first_name":      "", // skipped
		"handle":          "", // skipped
		"lowercase_value": "", // skipped
		"uppercase_value": "", // skipped
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/optional-string", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/optional-string", invalidPayload)

	fmt.Println("\n--- Skipped optional validation request ---")
	form.SendTestPost(":8080/optional-string", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
