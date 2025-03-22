package utility

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwtToken(username string) (string, error) {
	if secretKey == nil {
		return "", errors.New("JWT Secret Error")
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	//signed token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(username string, userId uint, userRole string) (string, error) {
	if secretKey == nil {
		return "", errors.New("JWT Secret Error")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": username,
		"userId":   userId,
		"userRole": userRole,
		"exp":      time.Now().Add(time.Hour * 5).Unix(),
	})

	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New("JWT Refresh Token Error")
	}
	return refreshToken, nil
}

func ValidateToken(jwtToken string) (*jwt.Token, jwt.MapClaims, error) {
	if secretKey == nil {
		return nil, nil, errors.New("JWT Secret Error")
	}

	//jwt.Parse(token, secretKey by checking signing methods)--> return decoded token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

		//signing method check
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims) //assert decoded token's claim with type jwt.MapClaims
	if !ok {
		return nil, nil, fmt.Errorf("invalid claims")
	}

	return token, claims, nil
}
