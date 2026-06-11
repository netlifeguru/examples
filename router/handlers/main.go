package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// Enable default middlewares for better logging, request IDs, etc.
	r.UseDefaults()

	// USING THE GENERAL FUNCTION: r.HandleFunc

	// Accepts absolutely all HTTP methods (ANY)
	r.HandleFunc("/any", "ANY", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "This catches anything!", "used_method": "%s"}`, req.Method)
	})

	// Accepts only selected methods (combination separated by space)
	r.HandleFunc("/documents", "GET POST", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "application/json")

		if req.Method == http.MethodGet {
			w.Write([]byte(`{"action": "Reading document list"}`))
		} else if req.Method == http.MethodPost {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"action": "Creating a new document"}`))
		}
	})

	// USING SPECIFIC SHORTCUTS FOR REST API

	// GET: Retrieve data (Read)
	r.GET("/users/{id}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"users": [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]}`))
	})

	// POST: Create a new record (Create)
	r.POST("/users", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "User successfully created", "id": 3}`))
	})

	// PUT: Complete replacement of an existing record (Update)
	// Notice the use of your custom parameter 'isDigits' for the ID
	r.PUT("/users/{id:isDigits}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		userID := ctx.Param("id")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "User %s has been completely replaced/updated"}`, userID)
	})

	// PATCH: Partial modification of an existing record (Partial Update)
	r.PATCH("/users/{id:isDigits}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		userID := ctx.Param("id")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "User %s record has been partially updated"}`, userID)
	})

	// DELETE: Remove a record (Delete)
	r.DELETE("/users/{id:isDigits}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		userID := ctx.Param("id")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "User %s has been deleted"}`, userID)
	})

	// SPECIAL HTTP METHODS

	// HEAD: Returns only headers, no body
	// The client (e.g., browser) uses this to check if a resource exists without downloading data.
	r.HEAD("/ping", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("X-System-Status", "Online")
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusOK)
		// With HEAD, the body is not sent to the client anyway
	})

	// OPTIONS: Discover allowed methods (used mainly for CORS preflight)
	r.OPTIONS("/api", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Allow", "OPTIONS, GET, POST, PUT, DELETE")
		w.WriteHeader(http.StatusNoContent) // 204 No Content
	})

	// TRACE: Diagnostic method - the server usually returns to the client exactly what it received
	r.TRACE("/echo", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "message/http")
		w.WriteHeader(http.StatusOK)

		// Simple echo of headers back to the client
		fmt.Fprintf(w, "%s %s %s\n", req.Method, req.URL.RequestURI(), req.Proto)
		for name, headers := range req.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	})

	// CONNECT: Establishing a network tunnel (used mainly with proxy servers)
	r.CONNECT("/tunnel", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Tunnel connection established (mock)"))
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
