package model

import "time"

type RankingNews struct{
	Id int	`gorm:"column:id;type:bigint;AUTO_INCREMENT;not null;primary_key"`
	Title string	`gorm:"column:title;type:varchar"`
	Content string	`gorm:"column:content"`
	Press string	`gorm:"column:press;type:varchar"`
	Date time.Time	`gorm:"column:date"`
	Url string	`gorm:"column:url"`
	Portal string	`gorm:"column:portal;type:varchar"`
	CreatedAt time.Time	`gorm:"column:created_at"`
	UpdatedAt time.Time	`gorm:"column:updated_at"`
}