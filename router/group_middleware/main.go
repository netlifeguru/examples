package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

// SimpleAuth is a mock authentication middleware
func SimpleAuth(next router.HandlerFunc) router.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "secret-key" {
			http.Error(w, "Unauthorized: Invalid API Key", http.StatusUnauthorized)
			return
		}
		next(w, r, ctx)
	}
}

func main() {
	rt := router.New()

	r := rt.With(func(next router.HandlerFunc) router.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
			fmt.Println(r.Header.Get("X-API-KEY"))
			next(w, r, ctx)
		}
	})

	// Global Middleware (Applies to EVERYTHING)
	r.Use(router.RequestID())

	// Public Group (No extra middleware)
	public := r.Group("/public")
	{
		public.GET("/info", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
			w.Write([]byte("This is public information."))
		})
	}

	// Protected Group (Applies SimpleAuth only to these routes)
	api := r.Group("/api")
	api.Use(SimpleAuth)
	{
		api.GET("/data", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
			w.Write([]byte("Secret data accessed successfully!"))
		})
	}

	// Resulting Routes:
	// GET /public/info -> RequestID
	// GET /api/data    -> RequestID + SimpleAuth

	if err := rt.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
