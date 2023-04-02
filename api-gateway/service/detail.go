package service

import (
	serviceModel "api-gateway/model"
	"context"
	"go.uber.org/zap"
)

type DetailRepository interface {
	Insert(ctx context.Context, detail serviceModel.Detail) (int64, error)
	Update(ctx context.Context, detail serviceModel.Detail) error
	GetOne(ctx context.Context, id int64) (*serviceModel.Detail, error)
	GetAll(ctx context.Context) ([]serviceModel.Detail, error)
	DeleteOne(ctx context.Context, id int64) error
}

type DetailService struct {
	logger           *zap.SugaredLogger
	detailRepository DetailRepository
}

func NewDetailService(logger *zap.SugaredLogger, detailRepository DetailRepository) DetailService {
	return DetailService{
		logger:           logger,
		detailRepository: detailRepository,
	}
}

func (s DetailService) AddDetails(ctx context.Context, details []serviceModel.Detail) []int64 {
	var ids []int64

	for _, detail := range details {
		id, err := s.detailRepository.Insert(ctx, detail)
		if err != nil {
			ids = append(ids, 0)
		} else {
			ids = append(ids, id)
		}
	}

	return ids
}

func (s DetailService) SelectAll(ctx context.Context) ([]serviceModel.Detail, error) {
	result, err := s.detailRepository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s DetailService) SelectById(ctx context.Context, id int64) (*serviceModel.Detail, error) {
	result, err := s.detailRepository.GetOne(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s DetailService) DeleteById(ctx context.Context, id int64) error {
	err := s.detailRepository.DeleteOne(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s DetailService) Update(ctx context.Context, detail serviceModel.Detail) error {
	err := s.detailRepository.Update(ctx, detail)

	if err != nil {
		return err
	}

	return nil
}
