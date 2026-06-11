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

	r.HandleFunc("/conditional-when", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ConditionalWhenRequest

		if !httpform.BindAndValidate(w, req, &in, ConditionalWhenSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("conditional when request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "conditional when validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"change_password":  true,
		"password":         "secret123", // validated because change_password is true
		"confirm_password": "secret123", // validated because change_password is true
		"account_type":     "company",
		"company_name":     "Acme s.r.o.",  // validated because account_type is company
		"vat_number":       "SK1234567890", // validated because account_type is company
	}

	invalidPayload := map[string]any{
		"change_password":  true,
		"password":         "short", // Required passes, MinLenWithCode fails
		"confirm_password": "",      // RequiredWithCode fails
		"account_type":     "company",
		"company_name":     "", // RequiredWithCode fails
		"vat_number":       "", // RequiredWithCode fails
	}

	skippedPayload := map[string]any{
		"change_password":  false,
		"password":         "", // not validated because change_password is false
		"confirm_password": "", // not validated because change_password is false
		"account_type":     "personal",
		"company_name":     "", // not validated because account_type is not company
		"vat_number":       "", // not validated because account_type is not company
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/conditional-when", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/conditional-when", invalidPayload)

	fmt.Println("\n--- Skipped conditional validation request ---")
	form.SendTestPost(":8080/conditional-when", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
