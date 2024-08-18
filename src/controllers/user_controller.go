package controllers

import (
	"api-customer/dtos"
	"api-customer/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userSrv services.UserService) *UserController {
	return &UserController{
		userService: userSrv,
	}
}

func (uc UserController) Pagination(c *gin.Context) {
	users := uc.userService.Pagination(c)
	c.JSON(http.StatusOK, users)
}

func (uc UserController) FindById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.userService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc UserController) Create(c *gin.Context) {
	var userCreateDTO dtos.UserCreateDTO
	err := c.ShouldBind(&userCreateDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = userCreateDTO.Validate()
	if err != nil {
		fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := uc.userService.Create(userCreateDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc UserController) Update(c *gin.Context) {
	var userUpdateDTO dtos.UserUpdateDTO
	err := c.ShouldBind(&userUpdateDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = userUpdateDTO.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	isUpdated := uc.userService.UpdateById(id, userUpdateDTO)
	if !isUpdated {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "updated")
}

func (uc UserController) Delete(c *gin.Context) {

	id := c.Param("id")
	isDeleted := uc.userService.DeleteById(id)
	if !isDeleted {
		c.JSON(http.StatusBadRequest, "user cannot delete")
		return
	}
	c.JSON(http.StatusOK, "deleted")
}
