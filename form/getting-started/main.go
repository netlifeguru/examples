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

	r.HandleFunc("/", "POST", func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		var in PostRequest

		if !httpform.BindAndValidate(w, r, &in, PostSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "request received",
			"data":    in,
		})
	})

	form.SendTestPost(":8080/", map[string]any{
		"name": "abcdefd",
		"age":  10,
	})

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
