package handlers

import (
	"encoding/json"
	"net/http"
)

// ExampleHandler shows the handler pattern. Replace with your domain handlers.
type ExampleHandler struct {
	// repo *repository.ExampleRepository
}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

func (h *ExampleHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "replace me"})
}
