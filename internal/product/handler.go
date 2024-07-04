package product

import (
	"context"
	"encoding/json"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/apperror"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/handlers"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

// NewHandler создаем структуру, но возвращаем интерфейс
func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

const (
	productsURL = "/products"
	productURL  = "/products/:uuid"
)

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, productsURL, apperror.Middleware(h.GetProducts))
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) error {
	all, err := h.repository.FindAll(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	marshal, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(500)
		logging.GetLogger().Error("Failed to marshal products")
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
	return nil
}
