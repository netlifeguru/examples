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

	r.HandleFunc("/custom-message", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in CustomMessageRequest

		opts := httpform.ResponseOptions{
			ValidationMessage:  "please check your input",
			InvalidJSONMessage: "request body is not valid JSON",
		}

		if !httpform.BindAndValidateWithOptions(w, req, &in, CustomMessageSchema(), 1<<20, opts) {
			fmt.Println("custom message validation failed")
			return
		}

		fmt.Println("custom message request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "custom message validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"email":    "john@example.com",
		"password": "secret123",
	}

	invalidPayload := map[string]any{
		"email":    "", // RequiredWithCode fails
		"password": "", // RequiredWithCode fails
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/custom-message", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/custom-message", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
