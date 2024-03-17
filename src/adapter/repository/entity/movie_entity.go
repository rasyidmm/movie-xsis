package entity

import "time"

type MovieEntity struct {
	Id          int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	Rating      float64    `gorm:"column:rating"`
	Image       string     `gorm:"column:image"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (MovieEntity) TableName() string {
	return "movie"
}
