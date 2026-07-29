package validator

import (
	"errors"
	"strings"

	"calculator/backend/internal/models"
)

// ValidateCalculateRequest performs basic validation on the request.
func ValidateCalculateRequest(req *models.CalculateRequest) error {
	if strings.TrimSpace(req.Operation) == "" {
		return errors.New("missing operation")
	}

	switch strings.ToLower(req.Operation) {
	case "add", "addition", "+", "subtract", "subtraction", "-", "multiply", "multiplication", "*", "x", "divide", "division", "/", "exponent", "exponentiation", "power", "^", "sqrt", "square root", "root", "√", "percentage", "percent", "%":
		return nil
	default:
		return errors.New("invalid operation")
	}
}
