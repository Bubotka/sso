package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/services/auth"
	"sso/internal/storage/sqlite"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.NewStorage(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.NewAuth(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.NewApp(log, authService, grpcPort)

	return &App{GRPCSrv: grpcApp}
}
