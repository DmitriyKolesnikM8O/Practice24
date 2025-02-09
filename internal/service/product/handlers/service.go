package handlers

import (
	_ "github.com/DmitriyKolesnikM8O/Practice24/docs"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/apperror"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/jwt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository"
	Handler "github.com/DmitriyKolesnikM8O/Practice24/internal/service"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/service/product/url"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

type service struct {
	logger     *logging.Logger
	repository repository.Repository
}

func NewSerivce(repository repository.Repository, logger *logging.Logger) Handler.Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (h *service) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, url.ProductsURL, apperror.Middleware(h.GetProducts))
	router.HandlerFunc(http.MethodGet, url.ProductURL, apperror.Middleware(h.GetProductByID))
	router.HandlerFunc(http.MethodPost, url.ProductsURL, jwt.JWTMiddleware(apperror.Middleware(h.CreateProduct)))
	router.HandlerFunc(http.MethodPut, url.ProductsURL, jwt.JWTMiddleware(apperror.Middleware(h.UpdateByID)))
	router.HandlerFunc(http.MethodGet, url.ReportURL, jwt.JWTMiddleware(apperror.Middleware(h.CreateReport)))
	router.HandlerFunc(http.MethodPost, url.AuthURL, apperror.Middleware(h.Auth))
	router.HandlerFunc(http.MethodDelete, url.DeleteURL, jwt.JWTMiddleware(apperror.Middleware(h.DeleteProduct)))
	router.HandlerFunc(http.MethodPost, url.RegisterURL, apperror.Middleware(h.UserRegister))
	router.HandlerFunc(http.MethodGet, "/swagger/*any", httpSwagger.Handler(httpSwagger.URL(url.SwaggerURL)))
}
