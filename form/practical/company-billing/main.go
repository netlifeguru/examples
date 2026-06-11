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

	r.HandleFunc("/company-billing", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in CompanyBillingRequest

		if !httpform.BindAndValidate(w, req, &in, CompanyBillingSchema(), 1<<20) {
			fmt.Println("company billing validation failed")
			return
		}

		fmt.Println("company billing request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "company billing validation passed",
			"data":    in,
		})
	})

	validCompanyPayload := map[string]any{
		"account_type":  "company",
		"company_name":  "Acme s.r.o.",
		"country":       "SK",
		"vat_number":    "SK1234567890",
		"billing_email": "billing@example.com",
	}

	validPersonalPayload := map[string]any{
		"account_type":  "personal",
		"company_name":  "", // not required for personal account
		"country":       "SK",
		"vat_number":    "", // not required for personal account
		"billing_email": "john@example.com",
	}

	invalidPayload := map[string]any{
		"account_type":  "company",
		"company_name":  "", // RequiredIfStrWithCode fails
		"country":       "SK",
		"vat_number":    "",              // RequiredIfStrWithCode fails
		"billing_email": "invalid-email", // EmailWithCode fails
	}

	fmt.Println("\n--- Valid company request ---")
	form.SendTestPost(":8080/company-billing", validCompanyPayload)

	fmt.Println("\n--- Valid personal request ---")
	form.SendTestPost(":8080/company-billing", validPersonalPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/company-billing", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
