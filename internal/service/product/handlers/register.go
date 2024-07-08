package handlers

import (
	"context"
	"encoding/json"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/auth"
	"golang.org/x/crypto/bcrypt"

	"net/http"
)

// @Summary register User
// @Tags auth
// @Description new user in table
// @ID register
// @Accept json
// @Produce json
// @Param user body auth.UserLogin true "Product information"
// @Success 200
// @Failure 400
// @Router /register [post]
func (h *service) UserRegister(w http.ResponseWriter, r *http.Request) error {
	var user auth.UserLogin
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("Failed to unmarshal user")
	}

	username, _ := h.repository.FindOneUser(context.Background(), user.Username)
	if username.Username != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already exists"))
		return nil
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	userForTable := auth.User{
		Username: user.Username,
		Password: password,
	}
	err := h.repository.CreateUser(context.Background(), &userForTable)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}
