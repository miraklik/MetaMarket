package app

import (
	grpcapp "internal/internal/app/grpc"
	"internal/internal/grpc/auth"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	gRPCapp := grpcapp.New(log, grpcPort)

	auth.Register(gRPCapp)

	return &App{
		GRPCServer: gRPCapp,
	}
}
