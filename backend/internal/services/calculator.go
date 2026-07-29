package services

import (
    "errors"
    "strings"
)

// CalculatorService defines calculator operations.
type CalculatorService interface {
    Calculate(op string, a, b float64) (float64, error)
}

type calculatorService struct{}

// NewCalculatorService constructs a new CalculatorService.
func NewCalculatorService() CalculatorService {
    return &calculatorService{}
}

// Calculate performs the requested operation on a and b.
func (s *calculatorService) Calculate(op string, a, b float64) (float64, error) {
    switch strings.ToLower(op) {
    case "add", "addition", "+":
        return a + b, nil
    case "subtract", "subtraction", "-":
        return a - b, nil
    case "multiply", "multiplication", "*", "x":
        return a * b, nil
    case "divide", "division", "/":
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
    default:
        return 0, errors.New("unknown operation")
    }
}
