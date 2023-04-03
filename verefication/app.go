package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"verefication/config"
	"verefication/database"
	"verefication/database/detailrepo"
	"verefication/lib/pctx"
	"verefication/service"
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
	a.logger.Debugf("server read %s  and write %s", a.settings.ReadTopic, a.settings.WriteTopic)
}

func (a App) Stop(ctx context.Context) {

	a.logger.Debugf("server stopped")
}
