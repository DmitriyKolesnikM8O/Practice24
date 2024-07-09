package app

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/config/config"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/repository/product"
	product2 "github.com/DmitriyKolesnikM8O/Practice24/internal/service/product/handlers"
	postgresql "github.com/DmitriyKolesnikM8O/Practice24/pkg/client/postgres"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

// Start starting app
func Start(router *httprouter.Router, cfg *config.Config) {
	var (
		listener  net.Listener
		ListenErr error
	)

	logger := logging.GetLogger()

	logger.Info("start app")
	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("listen unix socket")
		listener, ListenErr = net.Listen("unix", socketPath)
		logger.Infof("listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		address := fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
		listener, ListenErr = net.Listen("tcp", address)
		logger.Infof("listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}
	if ListenErr != nil {
		logger.Fatal(ListenErr)
	}

	logger.Info("Connect to postgres")
	postgreSQLClient, err := postgresql.NewClient(context.Background(), 3, cfg.Storage)
	if err != nil {
		logger.Fatal(err)
	}

	repository := product.NewRepository(postgreSQLClient, logger)
	logger.Info("register product service")
	productHandler := product2.NewSerivce(repository, logger)
	productHandler.Register(router)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Fatalln(server.Serve(listener))

}
