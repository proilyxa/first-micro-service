package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env-default:"0.0.0.0"`
	Port        string        `yaml:"port" env-default:"8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{
			Env: "dev",
			HTTPServer: HTTPServer{
				Host:        "0.0.0.0",
				Port:        "8082",
				Timeout:     time.Second * 4,
				IdleTimeout: time.Second * 60,
			},
		}
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &cfg
}
