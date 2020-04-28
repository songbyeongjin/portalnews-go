package model

import (
	"time"
)

type News struct{
	Id int	`gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	Title string	`gorm:"column:title"`
	Content string	`gorm:"column:content"`
	Press string	`gorm:"column:press"`
	Writer string	`gorm:"column:writer"`
	Date time.Time	`gorm:"column:date"`
	Url string	`gorm:"column:url"`
	Portal string	`gorm:"column:portal"`
	CreatedAt time.Time	`gorm:"column:created_at"`
	UpdatedAt time.Time	`gorm:"column:updated_at"`

}