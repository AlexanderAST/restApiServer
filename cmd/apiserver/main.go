package main

import (
	"flag"
	"log"

	"github.com/AlexanderAST/restApiServer/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal("Случилась ошибка с конфигом", err)
	}
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
