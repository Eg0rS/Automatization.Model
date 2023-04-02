package handler

import (
	"go.uber.org/zap"
	"io"
	"net/http"
)

func Ping(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if _, err := io.WriteString(w, "pong"); err != nil {
			logger.Errorf("error to write response: %s", err)
		}
	}
}
