package model

type User struct {
	UserId string `gorm:"column:user_id"`
	UserPass string `gorm:"column:user_pass"`
	Oauth string `gorm:"column:oauth"`
}