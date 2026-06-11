package main

import (
	"compress/gzip"
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// --- 1. Observability & Identification ---

	// Logger tracks request duration, method, and URL path.
	// Note: For detailed configuration and terminal output setup, see the Logging section.
	r.Use(router.Logger())

	// RealIP identifies the actual client IP, especially when behind a proxy.
	if err := router.SetTrustedProxies([]string{"10.0.0.0/8"}); err != nil {
		slog.Error("failed to set trusted proxies", "error", err)
	}

	r.Use(router.RealIP())

	// RequestID assigns a unique identifier to each request for tracing.
	r.Use(router.RequestID())

	// --- 2. Request Normalization ---

	// GetHead treats HEAD requests as GET requests for the underlying handlers.
	r.Use(router.GetHead())

	// CleanPath removes redundant slashes and normalizes the URL path.
	r.Use(router.CleanPath())

	// --- 3. Security & Headers ---

	// CORS handles Cross-Origin Resource Sharing and preflight OPTIONS requests.
	r.Use(router.CORS(router.CORSOptions{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// NoCache prevents browsers from caching the response.
	r.Use(router.NoCache())

	// --- 4. Optimization & Validation ---

	// Compress provides Gzip compression for specified content types.
	r.Use(router.Compress(
		gzip.DefaultCompression,
		"text/html",
		"text/plain",
		"text/css",
		"application/javascript",
		"text/javascript",
		"application/json",
	))

	// AllowContentType restricts requests based on the Content-Type header.
	r.Use(router.AllowContentType("application/json", "text/xml"))

	// ContentCharset validates the character encoding of the request.
	r.Use(router.ContentCharset([]string{"UTF-8", "Latin-1", ""}...))

	// Middleware
	r.Use(func(next router.HandlerFunc) router.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
			next(w, r, ctx)
		}
	})

	// --- 5. Routes & Handlers ---

	r.GET("/", func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`Hello World`))
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
