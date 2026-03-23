package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/yourorg/yourapp/middleware"
	"github.com/yourorg/yourapp/models"
	"github.com/yourorg/yourapp/repository"
)

// UserHandler handles user-related endpoints.
type UserHandler struct {
	users *repository.UserRepository
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(users *repository.UserRepository) *UserHandler {
	return &UserHandler{users: users}
}

// Me handles GET /api/me — returns the authenticated user's profile.
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.users.FindByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		slog.Error("me: find user", "error", err, "user_id", userID)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	})
}
