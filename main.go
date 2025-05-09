// main.go
package main

import (
	"fmt"
	"log/slog" // Structured logging (Go 1.21+)
	"net/http"
	"os"
	"time"
)

// loggingMiddleware logs request details.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request details before passing to the next handler
		slog.Info("request received",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
		)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log after the request has been handled
		slog.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", float64(time.Since(start).Microseconds())/1000.0, // Duration in milliseconds
			// Consider adding status code logging here if you have access to it
			// (requires a ResponseWriter wrapper to capture status)
		)
	})
}

// helloHandler handles requests to the /hello route.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Basic error handling for HTTP method
	if r.Method != http.MethodGet {
		slog.Warn("method not allowed for /hello", "method", r.Method, "path", r.URL.Path)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	slog.Debug("serving hello world response", "path", r.URL.Path)
	_, err := fmt.Fprint(w, "Hello, World from Go!")
	if err != nil {
		// This error is tricky because headers might have already been sent.
		// Log the error; the client might receive a partial response.
		slog.Error("error writing response to client", "error", err, "path", r.URL.Path)
	}
}

// rootHandler handles requests to the / route.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		slog.Warn("resource not found", "path", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	slog.Debug("serving root path response", "path", r.URL.Path)
	_, err := fmt.Fprint(w, "Welcome to the Go HTTP Server!")
	if err != nil {
		slog.Error("error writing response to client", "error", err, "path", r.URL.Path)
	}
}

func main() {
	// --- Structured Logger Setup ---
	// Choose between TextHandler (human-readable) or JSONHandler (machine-parseable)
	// var loggerHandler slog.Handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	var loggerHandler slog.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Log Debug, Info, Warn, Error levels
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Customize the "time" key to be "timestamp" and format it
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
				a.Value = slog.StringValue(a.Value.Time().Format(time.RFC3339Nano))
			}
			return a
		},
	})
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)

	// --- HTTP Server Setup ---
	mux := http.NewServeMux() // ServeMux is Go's standard HTTP request router

	// Register handlers
	mux.HandleFunc("/", rootHandler) // Handle requests to the root path
	mux.HandleFunc("/hello", helloHandler)

	// Wrap the main router with the logging middleware
	loggedMux := loggingMiddleware(mux)

	port := "8080"
	// Corrected line:
	slog.Info("Go HTTP server starting...", "port", port, "timestamp", time.Now().Format(time.RFC3339Nano))


	// Configure the server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      loggedMux, // Use the mux with logging middleware
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			slog.Info("Go HTTP server shut down gracefully.")
		} else {
			slog.Error("Go HTTP server failed to start or stopped unexpectedly", "error", err)
			os.Exit(1) // Exit if server fails to start
		}
	}
}