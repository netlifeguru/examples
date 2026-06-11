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

	r.HandleFunc("/conditional-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ConditionalRulesRequest

		if !httpform.BindAndValidate(w, req, &in, ConditionalRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("conditional rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "conditional rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"account_type":     "company",
		"company_name":     "Acme s.r.o.", // RequiredIfStr passes
		"role":             "user",
		"admin_note":       "",                 // ProhibitedIfStr passes
		"email":            "john@example.com", // RequiredWithAnyStr / RequiredWithoutAnyStr passes
		"phone":            "+421900123456",
		"contact_name":     "John Doe", // RequiredWithAnyStrWithCode passes
		"address":          "Main Street 1",
		"city":             "Bratislava",
		"country":          "SK",
		"vat_number":       "SK1234567890", // RequiredIfStrWithCode passes
		"password":         "secret123",    // When + MinLenWithCode passes
		"confirm_password": "secret123",    // When + Required passes
	}

	invalidPayload := map[string]any{
		"account_type":     "company",
		"company_name":     "", // RequiredIfStr fails
		"role":             "user",
		"admin_note":       "internal note", // ProhibitedIfStr fails
		"email":            "",              // RequiredWithAnyStr and RequiredWithoutAnyStrWithCode fail
		"phone":            "",              // RequiredWithoutAnyStr fails
		"contact_name":     "",              // RequiredWithAnyStrWithCode fails
		"address":          "Main Street 1",
		"city":             "",
		"country":          "SK",
		"vat_number":       "",      // RequiredIfStrWithCode fails
		"password":         "short", // When + MinLenWithCode fails
		"confirm_password": "",      // When + Required fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/conditional-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/conditional-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
