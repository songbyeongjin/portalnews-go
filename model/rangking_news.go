package model

import "time"

type RankingNews struct {
	Id              int       `gorm:"column:id;type:bigint;AUTO_INCREMENT;not null;primary_key"`
	Title           string    `gorm:"column:title;type:varchar"`
	TitleJapanese   string    `gorm:"column:title_ja;type:varchar"`
	Content         string    `gorm:"column:content"`
	ContentJapanese string    `gorm:"column:content_ja"`
	Press           string    `gorm:"column:press;type:varchar"`
	PressJapanese   string    `gorm:"column:press_ja"`
	Date            time.Time `gorm:"column:date"`
	Url             string    `gorm:"column:url"`
	Portal          string    `gorm:"column:portal;type:varchar"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
