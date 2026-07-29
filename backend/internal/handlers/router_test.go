package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockCalculatorForRouter struct{}

func (m *mockCalculatorForRouter) Calculate(op string, a, b float64) (float64, error) {
	return 0, nil
}

func TestNewRouter_RegistersCalculateRoute(t *testing.T) {
	logger := log.New(io.Discard, "", 0)
	router := NewRouter(&mockCalculatorForRouter{}, logger)

	req := httptest.NewRequest(http.MethodPost, "/calculate", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code == http.StatusNotFound {
		t.Fatalf("expected /calculate route to be registered, got 404")
	}
}
