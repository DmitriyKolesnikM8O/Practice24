package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/app/apiserver"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/app/config"

	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
