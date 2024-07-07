package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// @Summary GetProductByID
// @Tags products
// @Description one product from table by ID
// @ID get-product
// @Accept json
// @Param id path int true "Product id"
// @Success 200
// @Failure 400
// @Router /products/{id} [get]
func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.String()
	id = strings.TrimPrefix(id, "/products/")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("missing arguments: id")
		return nil
	}

	note, err := h.repository.FindOne(context.Background(), id)
	if err != nil || note.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	noteBytes, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("Failed to marshal product")
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(noteBytes)

	return nil
}
