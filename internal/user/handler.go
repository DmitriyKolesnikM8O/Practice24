package user

import (
	"github.com/DmitriyKolesnikM8O/Practice24/internal/handlers"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
	logger *logging.Logger
}

// NewHandler создаем структуру, но возвращаем интерфейс
func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

const (
	productsURL = "/products"
	productURL  = "/products/:uuid"
)

func (h *handler) Register(router *httprouter.Router) {
	router.GET(productsURL, h.GetProducts)
	router.POST(productsURL, h.CreateProduct)
	router.GET(productURL, h.GetProductByUUID)
	router.PUT(productURL, h.UpdateProduct)
	router.PATCH(productURL, h.PartiallyUpdateProduct)
	router.DELETE(productURL, h.DeleteProduct)
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("This is a page with list products"))
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("This is a page for create product"))
}

func (h *handler) GetProductByUUID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("This is a page for get product by uuid"))
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("This is a page for update product"))
}

func (h *handler) PartiallyUpdateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("This is a page for partially update product"))
}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("This is a page for delete product"))
}
