package handlers

import (
	"encoding/json"
	"net/http"
	"pathshala/internal/repository"
	"pathshala/internal/utils"
)

type AuthHandler struct {
	Repo *repository.AdminRepository
}

func NewAuthHandler(repo *repository.AdminRepository) *AuthHandler {
	return &AuthHandler{Repo: repo}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	admin, err := h.Repo.GetAdminByEmail(req.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	match := utils.CheckPasswordHash(req.Password, admin.Password)
	if !match {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(admin.ID, admin.Role, admin.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{
		Token: token,
		Role:  admin.Role,
	})
}

