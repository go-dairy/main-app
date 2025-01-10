package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `env:"ENVIRONMENT" env-default:"prod"`
	StorageHost string `env:"POSTGRES_HOST" env-required:"true"`
	StoragePort int    `env:"POSTGRES_PORT" env-default:"5432"`
	StorageUser string `env:"POSTGRES_USER" env-default:"postgres"`
	StoragePass string `env:"POSTGRES_PASS" env-default:"postgres"`
	StorageDB   string `env:"POSTGRES_DB" env-default:"postgres"`
	GRPC        GRPCConfig
}

type GRPCConfig struct {
	Port    int           `env:"GRPC_PORT" env-default:"50051"`
	Timeout time.Duration `env:"GRPC_TIMEOUT" env-default:"5"`
}

func MustLoad() *Config {

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
