package config

import (
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

// GetConfig Parse config from config.yml
func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()

		logger.Info("read application configuration")
		instance = &Config{}
		err := cleanenv.ReadConfig("config/config.yml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
