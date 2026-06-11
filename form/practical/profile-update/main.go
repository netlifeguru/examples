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

	r.HandleFunc("/profile-update", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ProfileUpdateRequest

		if !httpform.BindAndValidate(w, req, &in, ProfileUpdateSchema(), 1<<20) {
			fmt.Println("profile update validation failed")
			return
		}

		fmt.Println("profile update request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "profile update validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"nickname":   "johnny",
		"bio":        "I build Go APIs.",
		"website":    "https://example.com",
		"avatar_url": "https://example.com/avatar.png",
	}

	invalidPayload := map[string]any{
		"nickname":   "jo",                      // OptionalString runs because value is not empty; MinLenWithCode fails
		"bio":        string(make([]byte, 161)), // MaxLenWithCode fails
		"website":    "example.com",             // URLWithCode fails
		"avatar_url": "not-a-url",               // URLWithCode fails
	}

	skippedPayload := map[string]any{
		"nickname":   "", // OptionalString skipped
		"bio":        "", // OptionalString skipped
		"website":    "", // OptionalString skipped
		"avatar_url": "", // OptionalString skipped
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/profile-update", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/profile-update", invalidPayload)

	fmt.Println("\n--- Skipped optional validation request ---")
	form.SendTestPost(":8080/profile-update", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
