package middleware

import (
    "log"
    "net/http"
    "runtime/debug"
    "time"
)

// Logging returns a middleware that logs requests and durations.
func Logging(logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            logger.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
            next.ServeHTTP(w, r)
            logger.Printf("completed %s in %v", r.URL.Path, time.Since(start))
        })
    }
}

// Recover returns a middleware that recovers from panics and returns 500.
func Recover(logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            defer func() {
                if rec := recover(); rec != nil {
                    logger.Printf("panic: %v\n%s", rec, debug.Stack())
                    http.Error(w, "internal server error", http.StatusInternalServerError)
                }
            }()
            next.ServeHTTP(w, r)
        })
    }
}
