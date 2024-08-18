package services

import (
	"api-customer/dtos"
	"api-customer/models"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var key = []byte("jgxgfhghjikjo';kljk1223")

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (as AuthService) GetKey() []byte {
	return []byte(key)
}

func (as AuthService) GenerateToken(body dtos.AuthDTO) (resp dtos.AuthResponse, err error) {
	user, err := as.Login(body.UserName, body.Password)

	if err != nil {
		return resp, err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": user.UserName,
		"name":     user.Name,
		"id":       user.ID,
	})

	token, err := t.SignedString([]byte(key))

	if err != nil {
		return resp, err
	}

	resp.AccessToken = token
	return resp, nil
}

func (as AuthService) Login(userName, password string) (user models.User, err error) {
	err = as.db.Model(user).First(&user, "user_name = ?", userName).Error
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("username or password incorrect")
	}

	return user, nil
}

func ValidateToken(authorizationHeader string) (jwt.MapClaims, error) {
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
