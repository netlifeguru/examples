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

	r.HandleFunc("/response-map", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ResponseExampleRequest

		if !httpform.BindAndValidate(w, req, &in, ResponseExampleSchema(), 1<<20) {
			fmt.Println("map response validation failed")
			return
		}

		fmt.Println("map response request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "map response validation passed",
			"data":    in,
		})
	})

	r.HandleFunc("/response-flat", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ResponseExampleRequest

		if !httpform.BindAndValidateFlat(w, req, &in, ResponseExampleSchema(), 1<<20) {
			fmt.Println("flat response validation failed")
			return
		}

		fmt.Println("flat response request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "flat response validation passed",
			"data":    in,
		})
	})

	r.HandleFunc("/response-custom", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ResponseExampleRequest

		opts := httpform.ResponseOptions{
			ErrorFormat:        httpform.ErrorFormatMap,
			ValidationMessage:  "please check your input",
			InvalidJSONMessage: "request body is not valid JSON",
			UniqueCodes:        true,
		}

		if !httpform.BindAndValidateWithOptions(w, req, &in, ResponseExampleSchema(), 1<<20, opts) {
			fmt.Println("custom response validation failed")
			return
		}

		fmt.Println("custom response request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "custom response validation passed",
			"data":    in,
		})
	})

	r.HandleFunc("/response-custom-flat", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in ResponseExampleRequest

		opts := httpform.ResponseOptions{
			ErrorFormat:        httpform.ErrorFormatFlat,
			ValidationMessage:  "please check your input",
			InvalidJSONMessage: "request body is not valid JSON",
			UniqueCodes:        true,
		}

		if !httpform.BindAndValidateWithOptions(w, req, &in, ResponseExampleSchema(), 1<<20, opts) {
			fmt.Println("custom flat response validation failed")
			return
		}

		fmt.Println("custom flat response request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "custom flat response validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"email":    "john@example.com",
		"password": "secret123",
	}

	invalidPayload := map[string]any{
		"email":    "invalid-email",
		"password": "short",
	}

	emptyPayload := map[string]any{
		"email":    "",
		"password": "",
	}

	fmt.Println("\n--- Map response: valid request ---")
	form.SendTestPost(":8080/response-map", validPayload)

	fmt.Println("\n--- Map response: invalid request ---")
	form.SendTestPost(":8080/response-map", invalidPayload)

	fmt.Println("\n--- Flat response: invalid request ---")
	form.SendTestPost(":8080/response-flat", invalidPayload)

	fmt.Println("\n--- Custom map response: empty request ---")
	form.SendTestPost(":8080/response-custom", emptyPayload)

	fmt.Println("\n--- Custom flat response: empty request ---")
	form.SendTestPost(":8080/response-custom-flat", emptyPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
