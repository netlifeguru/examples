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

	r.HandleFunc("/sign-in", "POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		var in SignInRequest

		if !httpform.BindAndValidate(w, req, &in, SignInSchema(), 1<<20) {
			fmt.Println("form validation failed")
			return
		}

		fmt.Println("sign-in request received:", in)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "sign-in request received",
			"data":    in,
		})
	})

	form.SendTestPost(":8080/sign-in", SignInRequest{
		Email:    "john@example.com",
		Password: "secret123",
		Remember: false,
	})

	if err := r.ListenAndServe(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
