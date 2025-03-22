package model

type User struct {
	Abstract
	UserName string `json:"userName" gorm:"not null; size255;"`
	Password string `json:"password" gorm:"not null;"`
	Email    string `json:"email" gorm:"not null; size255; unique;"`
	Role     string `json:"role" gorm:"not null; size255; default:'user'"`
}

func (User) TableName() string {
	return "pp_users"
}
