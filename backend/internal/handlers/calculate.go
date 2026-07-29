package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"calculator/backend/internal/models"
	"calculator/backend/internal/services"
	v "calculator/backend/internal/validator"
)

func writeJSON(w http.ResponseWriter, status int, payload models.CalculateResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func calculateHandler(calc services.CalculatorService, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, models.CalculateResponse{Error: "method not allowed"})
			return
		}

		var req models.CalculateRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, models.CalculateResponse{Error: "invalid JSON payload"})
			return
		}

		if err := v.ValidateCalculateRequest(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, models.CalculateResponse{Error: err.Error()})
			return
		}

		result, err := calc.Calculate(req.Operation, req.A, req.B)
		if err != nil {
			if err.Error() == "division by zero" {
				writeJSON(w, http.StatusUnprocessableEntity, models.CalculateResponse{Error: err.Error()})
			} else {
				writeJSON(w, http.StatusBadRequest, models.CalculateResponse{Error: err.Error()})
			}
			return
		}

		writeJSON(w, http.StatusOK, models.CalculateResponse{Result: result})
	}
}
