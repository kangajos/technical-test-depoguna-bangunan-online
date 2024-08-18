package services

import (
	"api-customer/dtos"
	"api-customer/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us UserService) Pagination(c *gin.Context) paginate.Page {
	var user []models.User
	stmt := us.db.Model(user)
	if c.Query("name") != "" {
		stmt.Where("name like ?", "%"+c.Query("name")+"%")
	}
	if c.Query("userName") != "" {
		stmt.Where("user_name = ?", c.Query("userName"))
	}
	pg := paginate.New()
	paginate := pg.With(stmt).Request(c.Request).Response(&user)
	return paginate
}

func (us UserService) FindById(id string) (user models.User, err error) {
	ID, _ := strconv.ParseInt(id, 10, 64)
	err = us.db.Model(user).Preload("Orders").First(&user, "id = ?", ID).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("get user by id %d err: %s", id, err.Error()))
		return user, err
	}
	return user, nil
}

func (us UserService) Create(body dtos.UserCreateDTO) (user models.User, err error) {
	user.Name = body.Name
	user.UserName = body.UserName
	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	user.Password = string(passwordHashed)
	err = us.db.Save(&user).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("create user err: %s", err.Error()))
		return user, err
	}
	return user, nil
}

func (us UserService) UpdateById(id string, body dtos.UserUpdateDTO) bool {
	user, err := us.FindById(id)
	if err != nil {
		return false
	}
	user.Name = body.Name
	user.UserName = body.UserName
	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	user.Password = string(passwordHashed)
	err = us.db.Save(&user).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("update user err: %s", err.Error()))
		return false
	}

	return true
}

func (us UserService) DeleteById(id string) bool {
	ID, _ := strconv.ParseInt(id, 10, 64)
	err := us.db.Delete(&models.User{}, ID).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("delete user by id %d err: %s", id, err.Error()))
		return false
	}

	return true
}
