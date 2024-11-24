package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"orders-microservice/pkg/db/postgres"
	"orders-microservice/pkg/db/redis"
)

type Config struct {
	postgres.Config
	redis.RedisConfig

	GRPCServerPort int `env:"GRPC_SERVER_PORT" env-default:"50051"`
	RestServerPort int `env:"REST_SERVER_PORT" env-default:"8080"`
}

func New() *Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &cfg
}
