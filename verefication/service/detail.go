package service

import (
	"encoding/json"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"strconv"
	"time"
	"verefication/database/detailrepo"
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

	time.Sleep(4 * time.Second)
	var result detailrepo.DetailStageVersion

	result.DetailId = detail.Id
	if detail.Long > optimalLong || detail.Width > optimalWidth || detail.Height > optimalHeight || !slices.Contains(optimalColor, detail.Color) {
		result.Comment = "declined by invalid parameters"
		result.StageId = 2
		s.repository.Insert(result)
		return false, []byte{}, []byte{}
	}
	result.StageId = 3
	result.Comment = "passed verefication"
	s.repository.Insert(result)

	marshalDetail, _ := json.Marshal(detail)
	var i = int(detail.Id)

	return true, []byte(strconv.Itoa(i)), marshalDetail
}
