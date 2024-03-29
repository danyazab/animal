package main

import (
	"danyazab/animal/config"
	"danyazab/animal/internal"
	"danyazab/animal/internal/api"
	"log"
)

func main() {
	cfg := &config.Database{
		User:     "postgres",
		Password: "secret",
		Host:     "127.0.0.1",
		Port:     5432,
		Name:     "animal",
	}

	err := internal.Invoke(api.RunServer, cfg.Cgf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
