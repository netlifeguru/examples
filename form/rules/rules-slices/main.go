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

	r.HandleFunc("/slice-rules", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in SliceRulesRequest

		if !httpform.BindAndValidate(w, req, &in, SliceRulesSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("slice rules request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "slice rules validation passed",
			"data":    in,
		})
	})

	validPayload := map[string]any{
		"roles":             []string{"admin"},              // RequiredSliceStr
		"permissions":       []string{"read"},               // RequiredSliceStrWithCode
		"tags":              []string{"go", "validation"},   // MinItemsStr(2)
		"categories":        []string{"backend", "api"},     // MinItemsStrWithCode(2)
		"features":          []string{"login", "profile"},   // MaxItemsStr(3), DistinctStr
		"unique_codes":      []string{"ABC-123", "DEF-456"}, // MaxItemsStrWithCode(3), DistinctStrWithCode
		"required_items":    []string{"first", "second"},    // EachStr(VRequired)
		"min_length_items":  []string{"one", "two"},         // EachStr(VMinLen(3))
		"regex_items":       []string{"ABC-123", "DEF-456"}, // EachStr(VRegex)
		"custom_each_items": []string{"ABC-123", "DEF-456"}, // EachStr custom value rules
	}

	invalidPayload := map[string]any{
		"roles":             []string{},                                           // RequiredSliceStr fails
		"permissions":       []string{},                                           // RequiredSliceStrWithCode fails
		"tags":              []string{"go"},                                       // MinItemsStr fails
		"categories":        []string{"backend"},                                  // MinItemsStrWithCode fails
		"features":          []string{"login", "login", "admin", "x"},             // MaxItemsStr and DistinctStr fail
		"unique_codes":      []string{"ABC-123", "ABC-123", "DEF-456", "GHI-789"}, // MaxItemsStrWithCode and DistinctStrWithCode fail
		"required_items":    []string{"valid", ""},                                // required_items[1] fails
		"min_length_items":  []string{"ok", "valid"},                              // min_length_items[0] fails
		"regex_items":       []string{"ABC-123", "bad"},                           // regex_items[1] fails
		"custom_each_items": []string{"", "AB", "bad"},                            // multiple indexed errors
	}

	fmt.Println("\n--- Valid request ---")
	form.SendTestPost(":8080/slice-rules", validPayload)

	fmt.Println("\n--- Invalid request ---")
	form.SendTestPost(":8080/slice-rules", invalidPayload)

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
