package detailrepo

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"storage/database/detailrepo/query"
)

type Repository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewRepository(logger *zap.SugaredLogger, db *sqlx.DB) Repository {
	return Repository{
		logger: logger,
		db:     db,
	}
}

func (r Repository) Insert(detailStage DetailStageVersion) error {
	var id int64
	err := r.db.QueryRow(query.InsertDetailSql, detailStage.DetailId, detailStage.StageId, detailStage.Comment).Scan(&id)

	if err != nil {
		r.logger.Error(err)
	}

	return err
}

func (r Repository) GetOne(id int64) (Detail, error) {
	var result Detail

	err := r.db.Get(&result, query.SelectOneDetailSql, id)

	if err != nil {
		return Detail{}, err
	}

	return result, nil
}
