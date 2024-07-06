package product

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/DmitriyKolesnikM8O/Practice24/docs"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/auth"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/handlers"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/apperror"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/jwt"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger/v2"
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
	reportURL   = "/report"
	authURL     = "/auth"
	deleteUrl   = "/delete/:id"
)

// реализация интерфейса хэндлер
func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, productsURL, apperror.Middleware(h.GetProducts))
	router.HandlerFunc(http.MethodGet, productURL, apperror.Middleware(h.GetProductByID))
	router.HandlerFunc(http.MethodPost, productsURL, jwt.JWTMiddleware(apperror.Middleware(h.CreateProduct)))
	router.HandlerFunc(http.MethodPut, productsURL, jwt.JWTMiddleware(apperror.Middleware(h.UpdateByID)))
	router.HandlerFunc(http.MethodGet, reportURL, jwt.JWTMiddleware(apperror.Middleware(h.CreateReport)))
	router.HandlerFunc(http.MethodPost, authURL, apperror.Middleware(h.Auth))
	router.HandlerFunc(http.MethodDelete, deleteUrl, jwt.JWTMiddleware(apperror.Middleware(h.DeleteProduct)))

	router.HandlerFunc(http.MethodGet, "/swagger/*any", httpSwagger.Handler(
		httpSwagger.URL("http://0.0.0.0:1234/swagger/doc.json"), //The url pointing to API definition
	))
}

// @Summary GetProducts
// @Tags products
// @Description all products from table
// @ID get-product-by-id
// @Produce json
// @Success 200
// @Failure 400
// @Router /products [get]
func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) error {

	all, err := h.repository.FindAll(context.TODO())
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
	var dto CreateProduct
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
	note, err := h.repository.FindOne(context.TODO(), id)
	if err != nil {
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

// @Summary UodateProductByID
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
	var dto UpdateProduct
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
	all, sales, err := h.repository.FindAllForReport(context.TODO())
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

	marshalSales, err := json.Marshal(sales)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("Failed to marshal result sales")
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
	w.Write(marshalSales)

	return nil
}

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
func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.String()
	id = strings.TrimPrefix(id, "/delete/")

	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("missing arguments: id")
		return nil
	}
	err := h.repository.Delete(context.TODO(), id)
	if err != nil {
		h.logger.Error("Failed to delete product")
		w.WriteHeader(http.StatusInternalServerError)
		return err

	}

	w.WriteHeader(http.StatusOK)
	return nil
}
