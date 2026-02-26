package main

import (
	"net/http"
	"time"
)

func (app *application) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		app.logger.Info("request started",
			"method", r.Method,
			"path", r.URL.Path,
		)

		next.ServeHTTP(w, r)

		app.logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start).String(),
		)
	})
}
