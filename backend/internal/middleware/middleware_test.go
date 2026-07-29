package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCORS_OptionsRequest(t *testing.T) {
	handler := CORS()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodOptions, "/calculate", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", rr.Code)
	}
	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Fatalf("expected CORS allow origin *, got %q", rr.Header().Get("Access-Control-Allow-Origin"))
	}
}

func TestLogging_LogsRequest(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)
	handler := Logging(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/calculate", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(buf.String(), "GET /calculate") {
		t.Fatalf("expected log entry for request, got %q", buf.String())
	}
}

func TestRecover_PanicRecovered(t *testing.T) {
	handler := Recover(log.New(io.Discard, "", 0))(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "internal server error") {
		t.Fatalf("expected internal server error response, got %q", rr.Body.String())
	}
}
