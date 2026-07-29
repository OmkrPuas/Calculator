package services

import "testing"

func TestCalculatorService_Calculate(t *testing.T) {
    svc := NewCalculatorService()

    tests := []struct {
        name    string
        op      string
        a, b    float64
        want    float64
        wantErr bool
        errMsg  string
    }{
        {"add", "add", 5, 10, 15, false, ""},
        {"subtract", "subtract", 10, 4, 6, false, ""},
        {"multiply", "multiply", 3, 5, 15, false, ""},
        {"divide", "divide", 20, 4, 5, false, ""},
        {"divide_zero", "divide", 1, 0, 0, true, "division by zero"},
        {"unknown", "mod", 3, 2, 0, true, "unknown operation"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := svc.Calculate(tt.op, tt.a, tt.b)
            if tt.wantErr {
                if err == nil {
                    t.Fatalf("expected error but got nil")
                }
                if err.Error() != tt.errMsg {
                    t.Fatalf("expected error %q, got %q", tt.errMsg, err.Error())
                }
                return
            }
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            if got != tt.want {
                t.Fatalf("got %v want %v", got, tt.want)
            }
        })
    }
}
