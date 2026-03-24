package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yourorg/yourapp/handlers"
)

func TestExampleHandler_List(t *testing.T) {
	h := handlers.NewExampleHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/examples", nil)
	w := httptest.NewRecorder()

	h.List(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var body map[string]string
	if err := json.NewDecoder(w.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if _, ok := body["message"]; !ok {
		t.Error("expected 'message' field in response")
	}
}
