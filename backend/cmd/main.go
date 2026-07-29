package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"calculator/backend/internal/handlers"
	"calculator/backend/internal/middleware"
	"calculator/backend/internal/services"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	calcSvc := services.NewCalculatorService()
	router := buildRouter(calcSvc, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Println("starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server forced to shutdown: %v", err)
	}
	logger.Println("server stopped")
}

func buildRouter(calcSvc services.CalculatorService, logger *log.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/calculate", http.HandlerFunc(handlers.CalculateHandler(calcSvc, logger)))

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir != "" {
		logger.Printf("serving static assets from %s", staticDir)
		mux.Handle("/", staticHandler(staticDir))
	}

	handler := middleware.CORS()(mux)
	handler = middleware.Logging(logger)(handler)
	handler = middleware.Recover(logger)(handler)

	return handler
}

func staticHandler(root string) http.Handler {
	fs := http.FileServer(http.Dir(root))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Clean(r.URL.Path)
		if path == "/" || path == "" {
			http.ServeFile(w, r, filepath.Join(root, "index.html"))
			return
		}

		filePath := filepath.Join(root, path)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(root, "index.html"))
			return
		}

		fs.ServeHTTP(w, r)
	})
}
