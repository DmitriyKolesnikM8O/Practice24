package main

import (
	"github.com/DmitriyKolesnikM8O/Practice24/internal/app"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/config/config"
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

	app.Start(router, cfg)
}
