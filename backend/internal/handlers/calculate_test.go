package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
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

func TestCalculateHandler_Wrapper(t *testing.T) {
	reqBody, _ := json.Marshal(models.CalculateRequest{Operation: "add", A: 1, B: 2})
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := CalculateHandler(&mockCalculator{}, log.New(nil, "", 0))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp models.CalculateResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 3 {
		t.Fatalf("expected result 3, got %v", resp.Result)
	}
}

type mockCalculatorReturnError struct{}

func (m *mockCalculatorReturnError) Calculate(op string, a, b float64) (float64, error) {
	return 0, errors.New("division by zero")
}

func TestCalculateHandler_InvalidOperation(t *testing.T) {
	reqBody, _ := json.Marshal(models.CalculateRequest{Operation: "mod", A: 3, B: 2})
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(reqBody))
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

	if resp.Error != "invalid operation" {
		t.Fatalf("expected invalid operation error, got %q", resp.Error)
	}
}

func TestCalculateHandler_DivisionByZeroErrorBranch(t *testing.T) {
	reqBody, _ := json.Marshal(models.CalculateRequest{Operation: "divide", A: 10, B: 0})
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := calculateHandler(&mockCalculatorReturnError{}, log.New(nil, "", 0))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected status 422, got %d", rr.Code)
	}

	var resp models.CalculateResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Error != "division by zero" {
		t.Fatalf("expected division by zero error, got %q", resp.Error)
	}
}
