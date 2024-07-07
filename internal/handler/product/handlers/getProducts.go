package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

// @Summary GetProducts
// @Tags products
// @Description all products from table
// @ID get-product-by-id
// @Produce json
// @Success 200
// @Failure 400
// @Router /products [get]
func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) error {
	all, err := h.repository.FindAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	marshal, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("Failed to marshal products")
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshal)

	return nil
}
