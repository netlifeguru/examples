package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/netlifeguru/router"
)

// CustomRateLimitOpt demonstrates how to override the default rate limiter configuration.
// By default, the router calculates TTL and CleanupInterval automatically based on the threshold.
func CustomRateLimitOpt(cfg *router.RateLimitConfig) {
	// TTL: How long the client's tracking record is kept in memory.
	// We set it to 5 minutes to remember clients longer.
	cfg.TTL = 5 * time.Minute

	// CleanupInterval: How often the background garbage collector cleans up old records.
	// We set it to sweep every 1 minute.
	cfg.CleanupInterval = 1 * time.Minute
}

func main() {
	r := router.New()

	// Apply the RateLimit middleware globally.
	// 50 * time.Millisecond means a single client (IP + Route) can make at most 1 request every 50ms.
	// This equates to approximately 20 requests per second.
	// If the client exceeds this, the router automatically responds with HTTP 429 Too Many Requests.
	r.Use(router.RateLimit(50*time.Millisecond, CustomRateLimitOpt))

	// Simple GET endpoint to test the rate limiter
	r.HandleFunc("/", "GET", func(w http.ResponseWriter, req *http.Request, ctx *router.Context) {
		w.Write([]byte("Success! You are not rate-limited."))
	})

	if err := r.ListenAndServe(":8000"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
