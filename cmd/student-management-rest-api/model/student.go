package model

type Student struct {
	Abstract
	Name         string    `json:"name" gorm:"not null; size:250;"`
	Email        string    `json:"email" gorm:"not null; size:250; unique;"`
	Class        int       `json:"class" gorm:"not null; int type;"`
	RoleNo       int       `json:"roleNo" gorm:"not null; int type;"`
	Courses      []Course  `json:"courses" gorm:"many2many:pp_map_students_courses;"`
	FeesInvoices []Invoice `json:"fees_invoices" gorm:"foreignkey:StudentID"`
}

func (Student) TableName() string {
	return "pp_students"
}

//student --> invoice ---one2many
//student <--> course ---many2many
