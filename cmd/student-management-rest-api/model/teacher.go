package model

type Teacher struct {
	Abstract
	Name     string    `json:"name" gorm:"not null; size:250;"`
	Email    string    `json:"email" gorm:"not null; size:250; unique;"`
	Courses  []Course  `json:"courses" gorm:"many2many:pp_map_teachers_courses;"`
	Payslips []Invoice `json:"payslips" gorm:"foreignkey:TeacherID"`
}

func (Teacher) TableName() string {
	return "pp_teachers"
}

//teacher --> invoice ---one2many
//teacher <--> course ---many2many
