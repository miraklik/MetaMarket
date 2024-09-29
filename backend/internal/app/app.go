package app

import (
	grpcapp "internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	gRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		gRPCServer: grpcApp,
	}
}
