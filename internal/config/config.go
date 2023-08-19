package config

import (
	"time"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env-default:"0.0.0.0"`
	Port        uint16        `yaml:"port" env-default:"8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}
