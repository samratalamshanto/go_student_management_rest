package service

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/database"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/dto"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashSlice, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashSlice), nil
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(userReq *model.User) (*model.User, error) {
	hash, err := hashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}
	userReq.Password = hash

	res := database.DB.Create(userReq) //userReq is a pointer, so GORM updates it properly
	if res.Error != nil {
		return nil, res.Error
	}
	return userReq, nil
}

func Login(loginDto dto.LoginDto) (*model.User, error) {
	var user model.User
	res := database.DB.Where("user_name=?", loginDto.Username).Take(&user) //Take() does not use order by, First() uses order by

	if res.RowsAffected == 0 { //no record found
		return nil, fmt.Errorf("invalid username or password")
	}

	if res.Error != nil {
		return nil, res.Error
	}

	if verifyPassword(loginDto.Password, user.Password) {
		return &user, nil
	} else {
		return nil, fmt.Errorf("invalid username or password")
	}
}
