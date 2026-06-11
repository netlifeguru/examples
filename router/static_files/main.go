package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/netlifeguru/router"
)

func main() {
	r := router.New()

	// --- SERVING STATIC FILES & SMART FAVICON ---
	// The Static method maps a URL prefix to a local directory.
	// Example: r.Static("/static/", "./public")
	r.Static("/assets/", "./public")

	// --- THE "FAVICON BRIDGE" FEATURE ---
	// Modern browsers (like Firefox or Chrome) automatically look for 'favicon.ico'
	// at the root of your domain (http://localhost:8000/favicon.ico), regardless
	// of where your static assets are mapped.
	//
	// How NLG handles this:
	// 1. If 'favicon.ico' exists inside your static folder (e.g., ./public/favicon.ico),
	//    it is naturally available at http://localhost:8000/static/favicon.ico.
	// 2. The router AUTOMATICALLY registers an additional internal route at the root.
	// 3. This ensures that a request to http://localhost:8000/favicon.ico will
	//    NOT result in a 404 error, but will correctly serve the file from your folder.

	// Example Route:
	r.GET("/", func(w http.ResponseWriter, r *http.Request, ctx *router.Context) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<html>
				<head>
					<link rel="stylesheet" href="/assets/style.css">
				</head>
				<body>
					<h1>Static Files Example</h1>
					<p>Check the console to see if style.css and favicon.ico were loaded.</p>
				</body>
			</html>
		`))
	})

	// Start the server
	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
