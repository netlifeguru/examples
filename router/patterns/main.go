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

	// Only lowercase English letters (a-z)
	// Accepts: /tag/golang
	// Rejects: /tag/GoLang, /tag/go123
	r.GET("/tag/{slug:isLowerAlpha}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Lowercase tag matched: %s", slug)
	})

	// Only uppercase English letters (A-Z)
	// Accepts: /country/SK
	// Rejects: /country/Sk, /country/sk1
	r.GET("/country/{slug:isUpperAlpha}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Uppercase country code matched: %s", slug)
	})

	// Only English letters (a-z, A-Z)
	// Accepts: /name/JohnDoe
	// Rejects: /name/John-Doe, /name/John123
	r.GET("/name/{slug:isAlpha}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Alphabetic name matched: %s", slug)
	})

	// Only digits (0-9)
	// Accepts: /product/123456
	// Rejects: /product/123a, /product/12.34
	r.GET("/product/{slug:isDigits}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Product ID (digits only) matched: %s", slug)
	})

	// Alphanumeric characters only (a-z, A-Z, 0-9)
	// Accepts: /serial/A1B2C3
	// Rejects: /serial/A1-B2, /serial/A1_B2
	r.GET("/serial/{slug:isAlnum}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Serial number (alphanumeric) matched: %s", slug)
	})

	// Word characters (a-z, A-Z, 0-9, and underscore '_')
	// Accepts: /profile/john_doe_88
	// Rejects: /profile/john-doe
	r.GET("/profile/{slug:isWord}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Username matched: %s", slug)
	})

	// Article with a safe slug (letters, numbers, underscores, hyphens)
	// Accepts: /article/my_first-post-2024
	r.GET("/article/{slug:isSlugSafe}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Loading article with slug: %s", slug)
	})

	// Strict lowercase slug (only lowercase a-z, 0-9, and hyphens)
	// Accepts: /category/home-appliances
	// Rejects: /category/Home-Appliances, /category/home_appliances
	r.GET("/category/{slug:isSlug}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Strict lowercase category matched: %s", slug)
	})

	// Hexadecimal string (0-9, a-f, A-F)
	// Accepts: /color/ff00aa
	// Rejects: /color/ff00zz
	r.GET("/color/{slug:isHex}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Hex color matched: %s", slug)
	})

	// Secure path with UUID
	// Accepts: /user/123e4567-e89b-12d3-a456-426614174000
	r.GET("/user/{id:isUUID}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		userID := ctx.Param("id")
		fmt.Fprintf(w, "Valid UUID User ID: %s", userID)
	})

	// Safe text including spaces, dots, hyphens, and underscores
	// Accepts: /search/my log file v1.2-final_draft
	// Rejects: /search/my<script>log
	r.GET("/search/{slug:isSafeText}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Safe search text matched: %s", slug)
	})

	// Uppercase alphanumeric characters (A-Z, 0-9)
	// Accepts: /flight/FR1234
	// Rejects: /flight/fr1234
	r.GET("/flight/{slug:isUpperAlnum}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Flight number matched: %s", slug)
	})

	// Base64 encoded string
	// Accepts: /token/aGVsbG8=
	// Rejects: /token/aGVsbG8?
	r.GET("/token/{slug:isBase64}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Base64 token matched: %s", slug)
	})

	// Date string in YYYY-MM-DD format
	// Accepts: /events/2024-12-31
	// Rejects: /events/24-12-31, /events/2024/12/31
	r.GET("/events/{slug:isDateYMD}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Event date matched: %s", slug)
	})

	// 3. Downloading a file (isSafePath checks for slashes, dots, etc.)
	// Receives: /download/images/photo_01.jpg
	r.GET("/download/{filepath:isSafePath}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		filepath := ctx.Param("filepath")
		fmt.Fprintf(w, "Downloading safe file path: %s", filepath)
	})

	// Catch-all matcher, accepts anything except forward slashes ('/')
	// Accepts: /webhook/literally-anything-123!@#
	r.GET("/webhook/{slug:any}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		slug := ctx.Param("slug")
		fmt.Fprintf(w, "Catch-all webhook matched: %s", slug)
	})

	// Using your own regular expression directly in the URL
	// Only accepts exactly 4 digits (e.g. year)
	r.GET("/archive/{year:\\d{4}}", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		year := ctx.Param("year")
		fmt.Fprintf(w, "Archive for year: %s", year)
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
