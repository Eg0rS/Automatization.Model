package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"processing/config"
	"processing/database"
	"processing/database/detailrepo"
	"processing/lib/pctx"
	"processing/service"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	kafka    service.MyKafka
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}

	var (
		detailRepo    = detailrepo.NewRepository(logger, pgDb)
		detailService = service.NewDetailService(logger, detailRepo)
		kafka         = service.NewKafka(logger, settings, detailService)
	)

	return App{
		logger:   logger,
		settings: settings,
		kafka:    kafka,
	}
}

func (a App) Run() {
	go func() {
		fmt.Println("start consuming ... !!")
		a.kafka.Consume()
	}()
	a.logger.Debugf("server read %s ", a.settings.ReadTopic)
}

func (a App) Stop(ctx context.Context) {

	a.logger.Debugf("server stopped")
}
