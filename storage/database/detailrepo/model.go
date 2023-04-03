package detailrepo

import "time"

type Detail struct {
	Id        int64     `db:"id"`
	Long      float32   `db:"long"`
	Width     float32   `db:"width"`
	Height    float32   `db:"height"`
	Color     string    `db:"color"`
	EventDate time.Time `db:"event_date"`
	IsDeleted bool      `db:"is_deleted"`
}

type DetailStageVersion struct {
	Id        int64     `db:"id"`
	DetailId  int64     `db:"detail_id"`
	StageId   int64     `db:"stage_id"`
	Comment   string    `db:"comment"`
	EventDate time.Time `db:"event_date"`
}
