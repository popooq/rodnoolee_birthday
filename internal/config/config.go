package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address   string `env:"ADDRESS"`      // Address - адрес сервера
	DBAddress string `env:"DATABASE_DSN"` // DBAddress - адрес базы данных
}

func New() *Config {
	var (
		cfg Config
	)

	flag.StringVar(&cfg.Address, "a", "127.0.0.1:8080", "set server listening address")
	flag.StringVar(&cfg.DBAddress, "d", "mongodb://127.0.0.1:27017", "set the DB address")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		log.Printf("env parse failed :%s", err)
	}

	return &cfg
}
