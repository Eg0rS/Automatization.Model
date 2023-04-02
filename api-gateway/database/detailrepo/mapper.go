package detailrepo

import (
	"api-gateway/model"
)

func MapServiceToDb(detail model.Detail) Detail {
	return Detail{
		Id:        detail.Id,
		Long:      detail.Long,
		Width:     detail.Width,
		Height:    detail.Height,
		Color:     detail.Color,
		EventDate: detail.EventDate,
		IsDeleted: detail.IsDeleted,
	}
}

func MapListDbToService(dbDetails []Detail) []model.Detail {
	var result []model.Detail

	for _, dbDetail := range dbDetails {
		result = append(result, model.Detail{
			Id:        dbDetail.Id,
			Long:      dbDetail.Long,
			Width:     dbDetail.Width,
			Height:    dbDetail.Height,
			Color:     dbDetail.Color,
			EventDate: dbDetail.EventDate,
			IsDeleted: dbDetail.IsDeleted,
		})
	}

	return result
}

func MapDbToService(dbDetail Detail) model.Detail {

	return model.Detail{
		Id:        dbDetail.Id,
		Long:      dbDetail.Long,
		Width:     dbDetail.Width,
		Height:    dbDetail.Height,
		Color:     dbDetail.Color,
		EventDate: dbDetail.EventDate,
		IsDeleted: dbDetail.IsDeleted,
	}

}
