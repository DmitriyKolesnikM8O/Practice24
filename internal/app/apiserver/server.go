package apiserver

import (
	"github.com/DmitriyKolesnikM8O/Practice24/internal/app/config"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *config.Config
	logger *logrus.Logger
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting API server")
	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}
