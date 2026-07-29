package models

// CalculateRequest represents the JSON payload for calculation requests.
type CalculateRequest struct {
    Operation string  `json:"operation"`
    A         float64 `json:"a"`
    B         float64 `json:"b"`
}

// CalculateResponse represents a successful or error response.
type CalculateResponse struct {
    Result float64 `json:"result,omitempty"`
    Error  string  `json:"error,omitempty"`
}
