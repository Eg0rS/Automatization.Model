package handler

import (
	serviceModel "api-gateway/model"
	"api-gateway/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func AddDetail(logger *zap.SugaredLogger, service service.DetailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var details []serviceModel.Detail
		err := json.NewDecoder(r.Body).Decode(&details)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(service.AddDetails(r.Context(), details))
	}
}

func GetAllDetails(logger *zap.SugaredLogger, service service.DetailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		details, err := service.SelectAll(r.Context())
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(details)
	}
}

func GetDetailById(logger *zap.SugaredLogger, service service.DetailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		detail, err := service.SelectById(r.Context(), int64(id))
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(detail)
	}
}
func DeleteDetailById(logger *zap.SugaredLogger, service service.DetailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = service.DeleteById(r.Context(), int64(id))
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode("deleted")
	}
}

func UpdateDetail(logger *zap.SugaredLogger, service service.DetailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var detail serviceModel.Detail
		var err = json.NewDecoder(r.Body).Decode(&detail)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		if detail.Id == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("id is nil")
		}
		err = service.Update(r.Context(), detail)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode("updated")
	}
}
