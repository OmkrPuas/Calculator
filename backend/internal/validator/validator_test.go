package validator

import (
	"testing"

	"calculator/backend/internal/models"
)

func TestValidateCalculateRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     models.CalculateRequest
		wantErr bool
	}{
		{"valid_add", models.CalculateRequest{Operation: "add", A: 1, B: 2}, false},
		{"valid_sqrt", models.CalculateRequest{Operation: "sqrt", A: 9, B: 0}, false},
		{"invalid_op", models.CalculateRequest{Operation: "mod", A: 1, B: 2}, true},
		{"missing_op", models.CalculateRequest{Operation: "", A: 1, B: 2}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCalculateRequest(&tt.req)
			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
