package grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"orders-microservice/pkg/logger"
)

func ContextWithLogger(l logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		//l.Info(ctx, "request started", zap.String("method", info.FullMethod))
		resp, err = handler(ctx, req)
		if err != nil {
			l.Error(ctx, "request finished", zap.Error(err))
		}
		return resp, err
	}
}
