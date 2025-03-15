package service

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/database"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
)

func GetAllStudentData() (*[]model.Student, error) {
	var students []model.Student
	res := database.DB.Find(&students)
	return &students, res.Error
}

func GetStudentData(id uint) (*model.Student, error) {
	var student model.Student
	res := database.DB.First(&student, id)
	return &student, res.Error
}

func GetStudentDataEager(id uint) (*model.Student, error) {
	var student model.Student
	res := database.DB.Preload("Courses").Preload("FeesInvoices").First(&student, id)
	return &student, res.Error
}

func PostStudentData(student *model.Student) (*model.Student, error) {
	res := database.DB.Create(&student)
	return student, res.Error
}

func DeleteStudentData(id uint) error {
	res := database.DB.Unscoped().Delete(&model.Student{}, id)
	return res.Error
}
