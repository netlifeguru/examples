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

	r.HandleFunc("/change-password", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ChangePasswordRequest

		if !httpform.BindAndValidate(w, req, &in, ChangePasswordSchema(), 1<<20) {
			fmt.Println("change password validation failed")
			return
		}

		fmt.Println("change password request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "change password validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"change_password":  true,
		"current_password": "old-secret",
		"new_password":     "newsecret123",
		"confirm_password": "newsecret123",
	}

	invalidPayload := map[string]any{
		"change_password":  true,
		"current_password": "",      // RequiredWithCode fails
		"new_password":     "short", // MinLenWithCode fails
		"confirm_password": "",      // RequiredWithCode fails
	}

	skippedPayload := map[string]any{
		"change_password":  false,
		"current_password": "", // skipped
		"new_password":     "", // skipped
		"confirm_password": "", // skipped
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/change-password", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/change-password", invalidPayload)

	fmt.Println("\n--- Skipped conditional validation request ---")
	form.SendTestPost(":8080/change-password", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
