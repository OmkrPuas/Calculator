package handlers

import (
    "log"
    "net/http"

    "calculator/backend/internal/middleware"
    "calculator/backend/internal/services"
)

// NewRouter creates the HTTP router with middleware and handlers registered.
func NewRouter(calc services.CalculatorService, logger *log.Logger) http.Handler {
    mux := http.NewServeMux()
    mux.Handle("/calculate", http.HandlerFunc(calculateHandler(calc, logger)))

    // Apply middleware: logging then recovery
    handler := middleware.Logging(logger)(mux)
    handler = middleware.Recover(logger)(handler)

    return handler
}
