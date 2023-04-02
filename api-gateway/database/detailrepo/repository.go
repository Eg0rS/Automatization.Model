package detailrepo

import (
	"api-gateway/database/detailrepo/query"
	model "api-gateway/model"
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"log"
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

func (r Repository) Insert(ctx context.Context, detail model.Detail) (int64, error) {
	var data = MapServiceToDb(detail)
	var id int64
	err := r.db.QueryRowContext(ctx, query.InsertDetailSql, data.Long, data.Width, data.Height, data.Color).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (r Repository) Update(ctx context.Context, detail model.Detail) error {
	var result Detail

	err := r.db.GetContext(ctx, &result, query.SelectOneDetailSql, detail.Id)
	if err != nil {
		return err
	}

	if result.Long != detail.Long && detail.Long != nil {
		result.Long = detail.Long
	}
	if result.Width != detail.Width && detail.Width != nil {
		result.Width = detail.Width
	}
	if result.Height != detail.Height && detail.Height != nil {
		result.Height = detail.Height
	}
	if result.Color != detail.Color && detail.Color != nil {
		result.Color = detail.Color
	}
	if result.IsDeleted != detail.IsDeleted && detail.IsDeleted != nil {
		result.IsDeleted = detail.IsDeleted
	}

	_, err = r.db.ExecContext(ctx, query.UpdateDetailSql, result.Long, result.Width, result.Height, result.Color, result.IsDeleted, result.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetOne(ctx context.Context, id int64) (*model.Detail, error) {
	var result Detail

	err := r.db.GetContext(ctx, &result, query.SelectOneDetailSql, id)

	if err != nil {
		return nil, err
	}

	var r1 = MapDbToService(result)
	return &r1, nil
}

func (r Repository) GetAll(ctx context.Context) ([]model.Detail, error) {
	rows, err := r.db.QueryxContext(ctx, query.SelectAllDetailSql)

	var result []Detail

	for rows.Next() {
		detail := Detail{}
		err = rows.Scan(&detail.Id, &detail.Long, &detail.Width, &detail.Height, &detail.Color, &detail.EventDate, &detail.IsDeleted)
		if err != nil {
			log.Println(err)
		}
		result = append(result, detail)
	}

	if err != nil {
		return nil, err
	}
	return MapListDbToService(result), nil
}

func (r Repository) DeleteOne(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, query.DeleteDetailSql, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
