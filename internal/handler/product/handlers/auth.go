package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/handler/auth"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/jwt"
	"net/http"
)

// @Summary Auth
// @Tags auth
// @Description auth by username and password
// @ID auth
// @Accept json
// @Produce json
// @Param user body auth.User true "Username and password"
// @Success 200
// @Failure 400
// @Router /auth [post]
func (h *handler) Auth(w http.ResponseWriter, r *http.Request) error {

	var u auth.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error(err)
		return err
	}
	defer r.Body.Close()

	if u.Username != "admin" || u.Password != "admin" {
		h.logger.Error("Invalid username or password")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Invalid username or password"))
		return nil
	}

	token, err := jwt.GenerateJWT(u.Username)
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
