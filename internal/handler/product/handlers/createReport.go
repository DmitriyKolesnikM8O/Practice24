package handlers

import (
	"context"
	"encoding/json"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product/model"
	"net/http"
)

// @Summary Create Report
// @Security ApiKeyAuth
// @Tags products
// @Description Create report from table
// @ID report
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /report [get]
func (h *handler) CreateReport(w http.ResponseWriter, r *http.Request) error {
	salesProducts, sales, err := h.repository.FindAllForReport(context.Background())
	response := model.CombinedResponse{
		FirstType:  &salesProducts,
		SecondType: &sales,
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	marshalSales, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("Failed to marshal result sales")
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalSales)

	return nil
}
