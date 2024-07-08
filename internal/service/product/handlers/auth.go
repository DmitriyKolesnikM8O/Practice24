package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/jwt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/auth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// @Summary Auth
// @Tags auth
// @Description auth by username and password
// @ID auth
// @Accept json
// @Produce json
// @Param user body auth.UserLogin true "Username and password"
// @Success 200
// @Failure 400
// @Router /auth [post]
func (h *service) Auth(w http.ResponseWriter, r *http.Request) error {
	var user auth.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error(err)
		return err
	}

	defer r.Body.Close()
	username, _ := h.repository.FindOneUser(context.Background(), user.Username)
	if username.Username == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return nil
	}

	userFromTable, _ := h.repository.FindOneUser(context.Background(), username.Username)
	err := bcrypt.CompareHashAndPassword(userFromTable.Password, []byte(user.Password))
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Incorrect password"))
		return err
	}

	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		h.logger.Error("Failed to generate token")
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token":"%s"}`, token)))

	return nil
}
