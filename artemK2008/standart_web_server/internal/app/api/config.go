package api

import "github.com/artemK2008/standart_web_server/storage"

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    "8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
