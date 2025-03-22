package service

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/database"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
)

func GetAllCourseData() (*[]model.Course, error) {
	var courses []model.Course
	res := database.DB.Find(&courses)
	return &courses, res.Error
}

func GetCourseDataLazy(id uint) (*model.Course, error) {
	var course model.Course
	res := database.DB.First(&course, id)
	return &course, res.Error
}

func GetCourseDataEager(id uint) (*model.Course, error) {
	var course model.Course
	res := database.DB.
		Preload("Students", "active=? and class=?", true, 1).
		Preload("Teachers", "active=?", true).
		First(&course, id)
	return &course, res.Error
}

func PostCourseData(course *model.Course) (*model.Course, error) {
	res := database.DB.Create(&course)
	return course, res.Error
}

func DeleteCourseData(id uint) error {
	res := database.DB.Unscoped().Delete(&model.Course{}, id)
	return res.Error
}
