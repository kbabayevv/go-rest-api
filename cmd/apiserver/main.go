package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"http-rest-api/app/config"
	"http-rest-api/app/server"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".toml", "path to config file")
}

func main() {
	flag.Parse()
	newConfig := config.NewConfig()

	_, err := toml.DecodeFile(configPath, newConfig)

	if err != nil {
		log.Error("Couldn't read config file", err)
		return
	}

	s := server.NewAPIServer(newConfig)
	err = s.Start()
	if err != nil {
		log.Error("Couldn't start API server", err)
		return
	}
}
