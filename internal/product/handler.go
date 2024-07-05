package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/apperror"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/handlers"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
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
	productURL  = "/products/:id"
	report      = "/report"
)

// реализация интерфейса хэндлер
func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, productsURL, apperror.Middleware(h.GetProducts))
	router.HandlerFunc(http.MethodGet, productURL, apperror.Middleware(h.GetProductByID))
	router.HandlerFunc(http.MethodPost, productsURL, apperror.Middleware(h.CreateProduct))
	router.HandlerFunc(http.MethodPut, productsURL, apperror.Middleware(h.UpdateByID))
	router.HandlerFunc(http.MethodGet, report, apperror.Middleware(h.CreateReport))
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
		h.logger.Error("Failed to marshal products")
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
	return nil
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	var dto Product
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(400)
		h.logger.Error("Failed to unmarshal product")
	}

	id, err := h.repository.Create(context.TODO(), &dto)
	if err != nil {
		return err
	}
	w.Header().Set("ID number", fmt.Sprintf("%s", id))
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.String()
	id = strings.TrimPrefix(id, "/products/")

	if len(id) == 0 {
		w.WriteHeader(400)
		h.logger.Error("Failed to get product by id")
		return nil
	}
	note, err := h.repository.FindOne(context.TODO(), id)
	if err != nil {
		return err
	}
	noteBytes, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		h.logger.Error("Failed to marshal product")
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(noteBytes)
	return nil
}

func (h *handler) UpdateByID(w http.ResponseWriter, r *http.Request) error {
	var dto Product
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(400)
		h.logger.Error("Failed to unmarshal product for update")
	}
	err := h.repository.Update(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(500)
		h.logger.Error("Failed to update product")
		return err

	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *handler) CreateReport(w http.ResponseWriter, r *http.Request) error {
	all, sales, err := h.repository.FindAllForReport(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	marshal, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(500)
		h.logger.Error("Failed to marshal products")
		return err
	}

	marshalSales, err := json.Marshal(sales)
	if err != nil {
		w.WriteHeader(500)
		h.logger.Error("Failed to marshal result sales")
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
	w.Write(marshalSales)

	return nil
}
