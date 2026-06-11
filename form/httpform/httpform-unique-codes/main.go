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

	r.HandleFunc("/unique-codes-off", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in UniqueCodesRequest

		opts := httpform.ResponseOptions{
			ErrorFormat: httpform.ErrorFormatMap,
			UniqueCodes: false,
		}

		if !httpform.BindAndValidateWithOptions(w, req, &in, UniqueCodesSchema(), 1<<20, opts) {
			fmt.Println("unique codes off validation failed")
			return
		}

		fmt.Println("unique codes off request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "unique codes off validation passed",
			"data":    in,
		})
	})

	r.HandleFunc("/unique-codes-on", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in UniqueCodesRequest

		opts := httpform.ResponseOptions{
			ErrorFormat: httpform.ErrorFormatMap,
			UniqueCodes: true,
		}

		if !httpform.BindAndValidateWithOptions(w, req, &in, UniqueCodesSchema(), 1<<20, opts) {
			fmt.Println("unique codes on validation failed")
			return
		}

		fmt.Println("unique codes on request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "unique codes on validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"email": "john@example.com",
	}

	invalidPayload := map[string]any{
		"email": "", // RequiredWithCode fails twice
	}

	fmt.Println("\n--- UniqueCodes false: valid request ---")
	form.SendTestPost(":8080/unique-codes-off", validPayload)

	fmt.Println("\n--- UniqueCodes false: invalid request ---")
	form.SendTestPost(":8080/unique-codes-off", invalidPayload)

	fmt.Println("\n--- UniqueCodes true: invalid request ---")
	form.SendTestPost(":8080/unique-codes-on", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
