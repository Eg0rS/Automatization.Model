package service

import (
	"encoding/json"
	"go.uber.org/zap"
	"storage/database/detailrepo"
	"strconv"
	"time"
)

type DetailService struct {
	logger     *zap.SugaredLogger
	repository detailrepo.Repository
}

func NewDetailService(logger *zap.SugaredLogger, repository detailrepo.Repository) DetailService {
	return DetailService{
		logger:     logger,
		repository: repository,
	}
}

func (s DetailService) Processing(idDetail int64) (bool, []byte, []byte) {
	detail, err := s.repository.GetOne(idDetail)

	if err != nil {
		panic(err)
	}

	time.Sleep(100 * time.Millisecond)

	var result detailrepo.DetailStageVersion
	result.StageId = 2
	result.Comment = "passed storage"
	result.DetailId = detail.Id

	s.repository.Insert(result)

	marshalDetail, _ := json.Marshal(detail)
	var i = int(detail.Id)

	return true, []byte(strconv.Itoa(i)), marshalDetail
}
