package handlers

import (
	"context"
	"net/http"
	"strings"
)

// @Summary DeleteProduct
// @Security ApiKeyAuth
// @Tags products
// @Description delete product by ID
// @ID delete-product
// @Accept json
// @Param id path string true "Product id"
// @Success 200
// @Failure 400
// @Router /delete/{id} [delete]
func (h *service) DeleteProduct(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.String()
	id = strings.TrimPrefix(id, "/delete/")

	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("missing arguments: id")
		return nil
	}

	productFromTable, _ := h.repository.FindOne(context.Background(), id)
	if productFromTable.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		h.logger.Error("Product not found")
		return nil
	}

	err := h.repository.Delete(context.Background(), id)
	if err != nil {
		h.logger.Error("Failed to delete product")
		w.WriteHeader(http.StatusInternalServerError)
		return err

	}

	w.WriteHeader(http.StatusOK)

	return nil
}
