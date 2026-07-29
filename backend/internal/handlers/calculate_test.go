package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator/backend/internal/models"
)

type mockCalculator struct{}

func (m *mockCalculator) Calculate(op string, a, b float64) (float64, error) {
	return a + b, nil
}

func TestCalculateHandler_ReturnsResult(t *testing.T) {
	reqBody, _ := json.Marshal(models.CalculateRequest{Operation: "add", A: 3, B: 2})
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := calculateHandler(&mockCalculator{}, log.New(nil, "", 0))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp models.CalculateResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 5 {
		t.Fatalf("expected result 5, got %v", resp.Result)
	}
}

func TestCalculateHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader([]byte("{ invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := calculateHandler(&mockCalculator{}, log.New(nil, "", 0))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}

	var resp models.CalculateResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Error != "invalid JSON payload" {
		t.Fatalf("expected invalid JSON payload error, got %q", resp.Error)
	}
}
