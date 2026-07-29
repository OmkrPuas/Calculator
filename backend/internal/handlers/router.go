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

	handler := middleware.CORS()(mux)
	handler = middleware.Logging(logger)(handler)
	handler = middleware.Recover(logger)(handler)

	return handler
}
