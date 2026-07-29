package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "calculator/backend/internal/models"
    "calculator/backend/internal/services"
    v "calculator/backend/internal/validator"
)

func calculateHandler(calc services.CalculatorService, logger *log.Logger) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var req models.CalculateRequest
        dec := json.NewDecoder(r.Body)
        dec.DisallowUnknownFields()
        if err := dec.Decode(&req); err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            _ = json.NewEncoder(w).Encode(models.CalculateResponse{Error: "invalid JSON payload"})
            return
        }

        if err := v.ValidateCalculateRequest(&req); err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            _ = json.NewEncoder(w).Encode(models.CalculateResponse{Error: err.Error()})
            return
        }

        result, err := calc.Calculate(req.Operation, req.A, req.B)
        if err != nil {
            w.Header().Set("Content-Type", "application/json")
            // Domain errors like division by zero map to 422
            if err.Error() == "division by zero" {
                w.WriteHeader(http.StatusUnprocessableEntity)
            } else {
                w.WriteHeader(http.StatusBadRequest)
            }
            _ = json.NewEncoder(w).Encode(models.CalculateResponse{Error: err.Error()})
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _ = json.NewEncoder(w).Encode(models.CalculateResponse{Result: result})
    }
}
