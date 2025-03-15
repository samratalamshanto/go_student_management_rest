package model

type Course struct {
	Abstract
	CourseName  string    `json:"courseName" gorm:"size:100;not null"`
	CourseCode  string    `json:"courseCode" gorm:"size:100;not null;"`
	Description string    `json:"description" gorm:"size:500;"`
	Students    []Student `json:"students" gorm:"many2many:pp_map_students_courses;"`
	Teachers    []Teacher `json:"teachers" gorm:"many2many:pp_map_teachers_courses;"`
}

func (Course) TableName() string {
	return "pp_courses"
}

//teacher <--> course ---many2many
//student <--> course ---many2many
