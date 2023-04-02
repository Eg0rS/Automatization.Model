package api

import (
	"api-gateway/api/handler"
	"api-gateway/config"
	"api-gateway/lib/pctx"
	"api-gateway/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
)

func NewServer(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings, detailService service.DetailService) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.Ping(logger)).Methods(http.MethodGet)
	router.HandleFunc("/add", handler.AddDetail(logger, detailService)).Methods(http.MethodPost)
	router.HandleFunc("/get/{id}", handler.GetDetailById(logger, detailService)).Methods(http.MethodGet)
	router.HandleFunc("/get/all", handler.GetAllDetails(logger, detailService)).Methods(http.MethodGet)
	router.HandleFunc("/delete/{id}", handler.DeleteDetailById(logger, detailService)).Methods(http.MethodDelete)
	router.HandleFunc("/update", handler.UpdateDetail(logger, detailService)).Methods(http.MethodPatch)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
