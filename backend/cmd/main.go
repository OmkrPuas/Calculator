package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "calculator/backend/internal/handlers"
    "calculator/backend/internal/services"
)

func main() {
    logger := log.New(os.Stdout, "", log.LstdFlags)

    // Initialize services
    calcSvc := services.NewCalculatorService()

    // Initialize router (handlers will register routes)
    router := handlers.NewRouter(calcSvc, logger)

    srv := &http.Server{
        Addr:         ":8080",
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

    // Graceful shutdown on interrupt
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
