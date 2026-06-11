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

	r.HandleFunc("/string-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in StringRulesRequest

		if !httpform.BindAndValidate(w, req, &in, StringRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("string rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "string rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"username":         "johnny",            // MinLen(5)
		"nickname":         "tester",            // MinLenWithCode(5)
		"title":            "Developer",         // MaxLen(10)
		"short_title":      "Engineer",          // MaxLenWithCode(10)
		"pin":              "1234",              // Len(4)
		"access_code":      "ABC123",            // LenWithCode(6)
		"email":            "john@example.com",  // Email
		"contact_email":    "admin@example.com", // EmailWithCode
		"slug":             "hello-world-123",   // Regex
		"product_code":     "ABC-123",           // RegexWithCode
		"no_spaces":        "nospace",           // NotRegex whitespace
		"path":             "/users/profile",    // NotRegexWithCode admin
		"bio":              "I use form rules",  // Contains("form")
		"description":      "form package demo", // ContainsWithCode("package")
		"user_id":          "usr_123",           // StartsWith("usr_")
		"token":            "tok_abcdef",        // StartsWithWithCode("tok_")
		"image_file":       "avatar.png",        // EndsWith(".png")
		"document_file":    "contract.pdf",      // EndsWithWithCode(".pdf")
		"first_name":       "John",              // Alpha
		"last_name":        "Smith",             // AlphaWithCode
		"alphanum_code":    "User123",           // AlphaNum
		"reference":        "Ref123",            // AlphaNumWithCode
		"handle":           "user_name-123",     // AlphaDash
		"project_key":      "project_alpha-1",   // AlphaDashWithCode
		"lowercase_value":  "lowercase",         // Lowercase
		"lowercase_custom": "custom",            // LowercaseWithCode
		"uppercase_value":  "UPPERCASE",         // Uppercase
		"uppercase_custom": "CUSTOM",            // UppercaseWithCode
	}

	invalidPayload := map[string]any{
		"username":         "abc",              // MinLen fails
		"nickname":         "abc",              // MinLenWithCode fails
		"title":            "VeryLongTitle",    // MaxLen fails
		"short_title":      "VeryLongTitle",    // MaxLenWithCode fails
		"pin":              "123",              // Len fails
		"access_code":      "ABC12",            // LenWithCode fails
		"email":            "invalid-email",    // Email fails
		"contact_email":    "invalid-email",    // EmailWithCode fails
		"slug":             "Hello World!",     // Regex fails
		"product_code":     "abc-123",          // RegexWithCode fails
		"no_spaces":        "has space",        // NotRegex fails
		"path":             "/admin/settings",  // NotRegexWithCode fails
		"bio":              "I use validation", // Contains fails
		"description":      "form demo",        // ContainsWithCode fails
		"user_id":          "user_123",         // StartsWith fails
		"token":            "token_abcdef",     // StartsWithWithCode fails
		"image_file":       "avatar.jpg",       // EndsWith fails
		"document_file":    "contract.docx",    // EndsWithWithCode fails
		"first_name":       "John123",          // Alpha fails
		"last_name":        "Smith123",         // AlphaWithCode fails
		"alphanum_code":    "User-123",         // AlphaNum fails
		"reference":        "Ref-123",          // AlphaNumWithCode fails
		"handle":           "user name",        // AlphaDash fails
		"project_key":      "project.alpha",    // AlphaDashWithCode fails
		"lowercase_value":  "Lowercase",        // Lowercase fails
		"lowercase_custom": "Custom",           // LowercaseWithCode fails
		"uppercase_value":  "Uppercase",        // Uppercase fails
		"uppercase_custom": "Custom",           // UppercaseWithCode fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/string-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/string-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
