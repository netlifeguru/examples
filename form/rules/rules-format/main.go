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

	r.HandleFunc("/format-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in FormatRulesRequest

		if !httpform.BindAndValidate(w, req, &in, FormatRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("format rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "format rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"website":            "https://example.com",
		"callback_url":       "http://api.example.com/callback",
		"server_ip":          "127.0.0.1",
		"backup_ip":          "2001:db8::1",
		"resource_id":        "550e8400-e29b-41d4-a716-446655440000",
		"session_id":         "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"payload":            `{"name":"John","age":30}`,
		"metadata":           `["admin","editor"]`,
		"timezone":           "Europe/Bratislava",
		"preferred_timezone": "UTC",
	}

	invalidPayload := map[string]any{
		"website":            "example.com",
		"callback_url":       "ftp://example.com/callback",
		"server_ip":          "999.999.999.999",
		"backup_ip":          "not-an-ip",
		"resource_id":        "invalid-uuid",
		"session_id":         "123",
		"payload":            `{"name":"John",}`,
		"metadata":           `{invalid-json}`,
		"timezone":           "Europe/Invalid",
		"preferred_timezone": "Not/A_Timezone",
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/format-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/format-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
