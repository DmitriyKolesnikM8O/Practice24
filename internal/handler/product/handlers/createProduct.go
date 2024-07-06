package handlers

import (
	"context"
	"encoding/json"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"net/http"
)

// @Summary Create Product
// @Security ApiKeyAuth
// @Tags products
// @Description new product in table
// @ID create
// @Accept json
// @Produce json
// @Param product body CreateProduct true "Product information"
// @Success 200
// @Failure 400
// @Router /products [post]
func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	var dto model.CreateProduct
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("Failed to unmarshal product")
	}

	err := h.repository.Create(context.TODO(), &dto)
	if err != nil {
		return err
	}
	//w.Header().Set("ID number", fmt.Sprintf("%s", id))
	w.WriteHeader(http.StatusCreated)
	return nil
}
