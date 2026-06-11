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

	r.HandleFunc("/optional-slice", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in OptionalSliceRequest

		if !httpform.BindAndValidate(w, req, &in, OptionalSliceSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("optional slice request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "optional slice validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"tags":           []string{"go", "form"},         // OptionalSlice + MinItemsStrWithCode
		"permissions":    []string{"read", "write"},      // OptionalSlice + MaxItemsStrWithCode
		"features":       []string{"login", "profile"},   // OptionalSlice + DistinctStrWithCode
		"codes":          []string{"ABC-123", "DEF-456"}, // OptionalSlice + Min/Max/Distinct
		"required_items": []string{"first", "second"},    // OptionalSlice + EachStr(VRequiredWithCode)
		"min_len_items":  []string{"one", "two"},         // OptionalSlice + EachStr(VMinLenWithCode)
		"regex_items":    []string{"ABC-123", "DEF-456"}, // OptionalSlice + EachStr(VRegexWithCode)
		"custom_items":   []string{"ABC-123", "DEF-456"}, // OptionalSlice + EachStr custom rules
	}

	invalidPayload := map[string]any{
		"tags":           []string{"go"},                           // MinItemsStrWithCode fails
		"permissions":    []string{"read", "write", "delete", "x"}, // MaxItemsStrWithCode fails
		"features":       []string{"login", "login"},               // DistinctStrWithCode fails
		"codes":          []string{"ABC-123", "ABC-123"},           // DistinctStr fails
		"required_items": []string{"valid", ""},                    // required_items[1] fails
		"min_len_items":  []string{"ok", "valid"},                  // min_len_items[0] fails
		"regex_items":    []string{"ABC-123", "bad"},               // regex_items[1] fails
		"custom_items":   []string{"", "AB", "bad"},                // indexed validation errors
	}

	skippedPayload := map[string]any{
		"tags":           []string{}, // skipped
		"permissions":    []string{}, // skipped
		"features":       []string{}, // skipped
		"codes":          []string{}, // skipped
		"required_items": []string{}, // skipped
		"min_len_items":  []string{}, // skipped
		"regex_items":    []string{}, // skipped
		"custom_items":   []string{}, // skipped
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/optional-slice", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/optional-slice", invalidPayload)

	fmt.Println("\n--- Skipped optional validation request ---")
	form.SendTestPost(":8080/optional-slice", skippedPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
