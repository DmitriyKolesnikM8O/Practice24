package main

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/config"
	product2 "github.com/DmitriyKolesnikM8O/Practice24/internal/product"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/product/db"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/client/postgres"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("Connect to postgres")
	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatal(err)
	}
	repository := product.NewRepository(postgreSQLClient, logger)

	logger.Info("register product handler")
	productHandler := product2.NewHandler(repository, logger)
	productHandler.Register(router)

	//test Create new product
	//newProd := product2.Product{
	//	Name:  "помидор",
	//	Price: 24,
	//	Count: 3,
	//}
	//err = repository.Create(context.TODO(), &newProd)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//logger.Info("%v", newProd)

	//test FindOne
	//product, err := repository.FindOne(context.TODO(), "5")
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//logger.Info("%v", product)

	//connStr := "host=0.0.0.0 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//logger.Info("register handler")
	//handler := user.NewHandler(logger)
	//handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start app")

	var listener net.Listener
	var ListenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		//Abs возвращает абсолютный путь, Dir - все, кроме последнего элемента пути
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
		listener, ListenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if ListenErr != nil {
		logger.Fatal(ListenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatalln(server.Serve(listener))
}
