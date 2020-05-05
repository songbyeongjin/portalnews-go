package model

import "time"

type Review struct {
	ID      string    `gorm:"column:id;type:bigint;AUTO_INCREMENT;not null;primary_key"`
	NewsUrl string    `gorm:"column:news_url"`
	UserId  string    `gorm:"column:user_id"`
	Title   string    `gorm:"column:title"`
	Content string    `gorm:"column:content"`
	Date    time.Time `gorm:"column:date"`
}