package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator/backend/internal/models"
)

type mockCalculatorError struct{}

func (m *mockCalculatorError) Calculate(op string, a, b float64) (float64, error) {
	return 0, nil
}

func TestCalculateHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calculate", nil)
	rr := httptest.NewRecorder()

	handler := calculateHandler(&mockCalculatorError{}, log.New(nil, "", 0))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}

	var resp models.CalculateResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Error != "method not allowed" {
		t.Fatalf("expected method not allowed, got %q", resp.Error)
	}
}
