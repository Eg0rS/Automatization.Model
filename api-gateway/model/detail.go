package model

import "time"

type Detail struct {
	Id        *int64     `json:"id"`
	Long      *float32   `json:"long"`
	Width     *float32   `json:"width"`
	Height    *float32   `json:"height"`
	Color     *string    `json:"color"`
	EventDate *time.Time `json:"event_date"`
	IsDeleted *bool      `json:"is_deleted"`
}
