package config

import (
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
}

// сигнтон
var instance *Config

// примитив once, чтобы единожды распарсить
var once sync.Once

func GetConfig() *Config {
	//ровно 1 раз выполниться
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("reading config")
		instance = &Config{}

		//читаем и записываем конфиг
		err := cleanenv.ReadConfig("config.yml", instance)
		if err != nil {
			//выводим, что происходит не так
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
