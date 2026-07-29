package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type mockCalcService struct{}

func (m *mockCalcService) Calculate(op string, a, b float64) (float64, error) {
	return 0, nil
}

func TestBuildRouter_ServesCalculateRoute(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("STATIC_DIR", tmpDir)
	defer os.Unsetenv("STATIC_DIR")

	router := buildRouter(&mockCalcService{}, log.New(io.Discard, "", 0))
	req := httptest.NewRequest(http.MethodPost, "/calculate", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if rr.Code == http.StatusNotFound {
		t.Fatalf("expected /calculate route to be registered, got 404")
	}
}

func TestStaticHandler_ServesIndexForRoot(t *testing.T) {
	tmpDir := t.TempDir()
	indexFile := filepath.Join(tmpDir, "index.html")
	if err := os.WriteFile(indexFile, []byte("<html>ok</html>"), 0644); err != nil {
		t.Fatal(err)
	}

	handler := staticHandler(tmpDir)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if got := rr.Body.String(); got != "<html>ok</html>" {
		t.Fatalf("expected index body, got %q", got)
	}
}

func TestStaticHandler_FallsBackToIndexForMissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	indexFile := filepath.Join(tmpDir, "index.html")
	if err := os.WriteFile(indexFile, []byte("<html>fallback</html>"), 0644); err != nil {
		t.Fatal(err)
	}

	handler := staticHandler(tmpDir)
	req := httptest.NewRequest(http.MethodGet, "/missing.js", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if got := rr.Body.String(); got != "<html>fallback</html>" {
		t.Fatalf("expected fallback index body, got %q", got)
	}
}
