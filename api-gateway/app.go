package main

import (
	"api-gateway/api"
	"api-gateway/config"
	"api-gateway/database"
	"api-gateway/database/detailrepo"
	"api-gateway/lib/pctx"
	"api-gateway/service"
	"context"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}

	err = database.UpMigrations(pgDb)
	if err != nil {
		panic(err)
	}

	var (
		detailRepo    = detailrepo.NewRepository(logger, pgDb)
		detailService = service.NewDetailService(logger, detailRepo)
		server        = api.NewServer(ctxProvider, logger, settings, detailService)
	)

	return App{
		logger:   logger,
		settings: settings,
		server:   server,
	}
}

func (a App) Run() {
	go func() {
		_ = a.server.ListenAndServe()
	}()
	a.logger.Debugf("HTTP server started on %d", a.settings.Port)
}

func (a App) Stop(ctx context.Context) {
	_ = a.server.Shutdown(ctx)
	a.logger.Debugf("HTTP server stopped")
}
