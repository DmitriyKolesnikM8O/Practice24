package main

import (
	"context"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/app"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/config/config"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product"
	product2 "github.com/DmitriyKolesnikM8O/Practice24/internal/service/product/handlers"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/client/postgres"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// @Title Generating Report API
// @version 1.0
// @description API service to generate report based on monthly sales

// @host 0.0.0.0:1234
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()
	cfg := config.GetConfig()

	logger.Info("Connect to postgres")
	postgreSQLClient, err := postgresql.NewClient(context.Background(), 3, cfg.Storage)
	if err != nil {
		logger.Fatal(err)
	}

	repository := product.NewRepository(postgreSQLClient, logger)
	logger.Info("register product service")
	productHandler := product2.NewSerivce(repository, logger)
	productHandler.Register(router)

	app.Start(router, cfg)
}
