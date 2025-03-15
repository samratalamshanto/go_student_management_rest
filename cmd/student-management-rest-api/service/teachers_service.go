package service

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/database"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
)

func GetAllTeacherData() (*[]model.Teacher, error) {
	var teachers []model.Teacher
	res := database.DB.Find(&teachers)
	return &teachers, res.Error
}

func GetTeacherData(id uint) (*model.Teacher, error) {
	var teacher model.Teacher
	res := database.DB.First(&teacher, id)
	return &teacher, res.Error
}

func GetTeacherDataEager(id uint) (*model.Teacher, error) {
	var teacher model.Teacher
	res := database.DB.Preload("Courses").Preload("Payslips").First(&teacher, id)
	return &teacher, res.Error
}

func PostTeacherData(teacher *model.Teacher) (*model.Teacher, error) {
	res := database.DB.Create(&teacher)
	return teacher, res.Error
}

func DeleteTeacherData(id uint) error {
	res := database.DB.Unscoped().Delete(&model.Teacher{}, id)
	return res.Error
}
