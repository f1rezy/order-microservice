package main

import (
	"context"
	"fmt"
	"orders-microservice/internal/config"
	"orders-microservice/internal/repository"
	service "orders-microservice/internal/service"
	"orders-microservice/internal/transport/grpc"
	"orders-microservice/pkg/db/postgres"
	"orders-microservice/pkg/db/redis"
	"orders-microservice/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

const (
	serviceName = "orders"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	cfg := config.New()
	if cfg == nil {
		panic("failed to load config")
	}

	db, err := postgres.New(cfg.Config)
	if err != nil {
		mainLogger.Error(ctx, err.Error())
		return
	}

	cache := redis.New(cfg.RedisConfig)
	fmt.Println(cache.Ping(ctx))

	repo := repository.NewOrderRepository(db)

	srv := service.NewOrderService(repo)

	grpcserver, err := grpc.New(ctx, cfg.GRPCServerPort, cfg.RestServerPort, srv)
	if err != nil {
		mainLogger.Error(ctx, err.Error())
		return
	}

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcserver.Start(ctx); err != nil {
			mainLogger.Error(ctx, err.Error())
		}
	}()

	<-graceCh

	mainLogger.Info(ctx, "Server Stopped")
}
