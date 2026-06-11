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

	r.HandleFunc("/tags", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in TagsRequest

		if !httpform.BindAndValidate(w, req, &in, TagsSchema(), 1<<20) {
			fmt.Println("tags validation failed")
			return
		}

		fmt.Println("tags request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "tags validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"tags": []string{"go", "form-package", "validation"},
	}

	invalidPayload := map[string]any{
		"tags": []string{
			"go",
			"go",          // DistinctStrWithCode fails
			"",            // VRequiredWithCode fails
			"Invalid Tag", // VRegexWithCode fails
			"api",
			"backend", // MaxItemsStrWithCode fails
		},
	}

	emptyPayload := map[string]any{
		"tags": []string{}, // RequiredSliceStrWithCode and MinItemsStrWithCode fail
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/tags", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/tags", invalidPayload)

	fmt.Println("\n--- Empty request ---")
	form.SendTestPost(":8080/tags", emptyPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
