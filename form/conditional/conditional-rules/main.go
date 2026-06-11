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
		"account_type":  "company",
		"company_name":  "Acme s.r.o.", // RequiredIfStr passes
		"country":       "SK",
		"vat_number":    "SK1234567890", // RequiredIfStrWithCode passes
		"role":          "user",
		"admin_note":    "", // ProhibitedIfStr passes
		"user_type":     "external",
		"internal_note": "", // ProhibitedIfStrWithCode passes
		"phone":         "+421900123456",
		"email":         "john@example.com", // RequiredWithAnyStr passes
		"address":       "Main Street 1",
		"city":          "",
		"contact_name":  "John Doe", // RequiredWithAnyStrWithCode passes
		"backup_email":  "backup@example.com",
		"backup_phone":  "", // RequiredWithoutAnyStrWithCode passes because backup_email is filled
	}

	invalidPayload := map[string]any{
		"account_type":  "company",
		"company_name":  "", // RequiredIfStr fails
		"country":       "SK",
		"vat_number":    "", // RequiredIfStrWithCode fails
		"role":          "user",
		"admin_note":    "internal only", // ProhibitedIfStr fails
		"user_type":     "external",
		"internal_note": "secret note", // ProhibitedIfStrWithCode fails
		"phone":         "+421900123456",
		"email":         "", // RequiredWithAnyStr fails
		"address":       "Main Street 1",
		"city":          "",
		"contact_name":  "", // RequiredWithAnyStrWithCode fails
		"backup_email":  "", // RequiredWithoutAnyStrWithCode fails
		"backup_phone":  "",
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
