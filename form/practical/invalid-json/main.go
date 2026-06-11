package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/form"
	"github.com/netlifeguru/form/httpform"
	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	r.HandleFunc("/invalid-json", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in InvalidJSONRequest

		opts := httpform.ResponseOptions{
			InvalidJSONMessage: "request body is not valid JSON",
		}

		if !httpform.BindAndValidateWithOptions(w, req, &in, InvalidJSONSchema(), 1<<20, opts) {
			fmt.Println("invalid json request failed")
			return
		}

		fmt.Println("invalid json request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "invalid json example validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"email": "john@example.com",
	}

	validationPayload := map[string]any{
		"email": "invalid-email", // valid JSON, but validation fails
	}

	invalidJSON := `{"email": "john@example.com"`

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/invalid-json", validPayload)

	fmt.Println("\n--- Validation error request ---")
	form.SendTestPost(":8080/invalid-json", validationPayload)

	fmt.Println("\n--- Invalid JSON request ---")
	sendRawPost("http://localhost:8080/invalid-json", invalidJSON)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

func sendRawPost(url string, body string) {
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(body))
	if err != nil {
		fmt.Println("--- Raw Client: request failed:", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("--- Raw Client: failed to read response:", err)
		return
	}

	fmt.Printf("--- Raw Client: Status %s, Response: %s\n", resp.Status, string(data))
}
