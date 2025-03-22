package model

type User struct {
	Abstract
	UserName string `json:"userName" gorm:"not null; size255;"`
	Password string `json:"password" gorm:"not null;"`
	Email    string `json:"email" gorm:"not null; size255; unique;"`
}

func (User) TableName() string {
	return "pp_users"
}
