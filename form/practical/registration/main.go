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

	r.HandleFunc("/registration", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in RegistrationRequest

		if !httpform.BindAndValidate(w, req, &in, RegistrationSchema(), 1<<20) {
			fmt.Println("registration validation failed")
			return
		}

		fmt.Println("registration request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "registration validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"name":             "John",
		"email":            "john@example.com",
		"password":         "secret123",
		"password_confirm": "secret123",
		"age":              18,
		"accept_terms":     true,
	}

	invalidPayload := map[string]any{
		"name":             "J",             // MinLen(2) fails
		"email":            "invalid-email", // EmailWithCode fails
		"password":         "short",         // MinLenWithCode(8) fails
		"password_confirm": "",              // Required fails
		"age":              17,              // Min(18) fails
		"accept_terms":     false,           // IsTrueWithCode fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/registration", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/registration", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
