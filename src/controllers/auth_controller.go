package controllers

import (
	"api-customer/dtos"
	"api-customer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authSrv services.AuthService
}

func NewAuthController(authSrv services.AuthService) *AuthController {
	return &AuthController{
		authSrv: authSrv,
	}
}

func (ac AuthController) Login(c *gin.Context) {
	var body dtos.AuthDTO

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := ac.authSrv.GenerateToken(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (ac AuthController) Logout(c *gin.Context) {
	var body dtos.AuthDTO

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, body)
}
