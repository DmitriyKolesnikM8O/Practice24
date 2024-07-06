package handlers

import (
	"context"
	"encoding/json"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"net/http"
)

// @Summary UpdateProductByID
// @Security ApiKeyAuth
// @Tags products
// @Description update product by ID
// @ID update-product
// @Accept json
// @Param information body UpdateProduct true "Product information"
// @Success 200
// @Failure 400
// @Router /products [put]
func (h *handler) UpdateByID(w http.ResponseWriter, r *http.Request) error {
	var dto model.UpdateProduct
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("Failed to unmarshal product for update")
	}
	err := h.repository.Update(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("Failed to update product")
		return err

	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
