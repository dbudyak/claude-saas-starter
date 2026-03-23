package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/yourorg/yourapp/config"
	"github.com/yourorg/yourapp/models"
	"github.com/yourorg/yourapp/repository"
)

// AuthHandler handles authentication endpoints.
type AuthHandler struct {
	users *repository.UserRepository
	cfg   *config.Config
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(users *repository.UserRepository, cfg *config.Config) *AuthHandler {
	return &AuthHandler{users: users, cfg: cfg}
}

// Signup handles POST /api/auth/signup
func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if req.Email == "" || !strings.Contains(req.Email, "@") {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 8 {
		http.Error(w, "password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	exists, err := h.users.EmailExists(r.Context(), req.Email)
	if err != nil {
		slog.Error("signup: check email exists", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "email already registered", http.StatusConflict)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("signup: hash password", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	user, err := h.users.Create(r.Context(), req.Email, string(hash))
	if err != nil {
		slog.Error("signup: create user", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	token, err := h.issueToken(user.ID)
	if err != nil {
		slog.Error("signup: issue token", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.LoginResponse{
		Token: token,
		User: models.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	})
}

// Login handles POST /api/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	user, err := h.users.FindByEmail(r.Context(), req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			// Generic message — don't reveal whether user exists
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		slog.Error("login: find user", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.issueToken(user.ID)
	if err != nil {
		slog.Error("login: issue token", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.LoginResponse{
		Token: token,
		User: models.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	})
}

// issueToken creates a signed JWT for the given user ID.
func (h *AuthHandler) issueToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.cfg.JWTSecret))
}
