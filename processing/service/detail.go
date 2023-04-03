package service

import (
	"encoding/json"
	"go.uber.org/zap"
	"processing/database/detailrepo"
	"strconv"
	"time"
)

var optimalLong = 15.0
var optimalWidth = 15.0
var optimalHeight = 15.0
var optimalColor = []string{"black", "white", "green", "purple"}

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

	time.Sleep(8 * time.Second)
	var result detailrepo.DetailStageVersion
	result.StageId = 4
	result.DetailId = detail.Id
	result.Comment = "passed processing"

	s.repository.Insert(result)

	marshalDetail, _ := json.Marshal(detail)
	var i = int(detail.Id)

	return true, []byte(strconv.Itoa(i)), marshalDetail
}
