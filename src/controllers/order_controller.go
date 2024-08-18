package controllers

import (
	"api-customer/dtos"
	"api-customer/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderSrv services.OrderService) *OrderController {
	return &OrderController{
		orderService: orderSrv,
	}
}

func (uc OrderController) Pagination(c *gin.Context) {
	bodyUserId, _ := c.Get("userId")

	// Handle the case where bodyUserId is a float64
	userIdFloat, ok := bodyUserId.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("userId is not a valid number"))
		return
	}

	// Convert float64 to int64
	userId := int64(userIdFloat)
	orders := uc.orderService.Pagination(c, userId)
	c.JSON(http.StatusOK, orders)
}

func (uc OrderController) FindById(c *gin.Context) {
	id := c.Param("id")
	order, err := uc.orderService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, order)
}

func (uc OrderController) Create(c *gin.Context) {
	var orderCreateDTO dtos.OrderCreateDTO
	err := c.ShouldBind(&orderCreateDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse(err.Error()))
		return
	}
	bodyUserId, _ := c.Get("userId")

	// Handle the case where bodyUserId is a float64
	userIdFloat, ok := bodyUserId.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("userId is not a valid number"))
		return
	}

	// Convert float64 to int64
	userId := int64(userIdFloat)
	orderCreateDTO.UserID = userId
	err = orderCreateDTO.Validate()
	if err != nil {
		fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	order, err := uc.orderService.Create(orderCreateDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, dtos.SuccessResponse(order, "Data has been Created"))
}

func (uc OrderController) Update(c *gin.Context) {
	var orderUpdateDTO dtos.OrderUpdateDTO
	err := c.ShouldBind(&orderUpdateDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse(err.Error()))
		return
	}

	bodyUserId, _ := c.Get("userId")

	// Handle the case where bodyUserId is a float64
	userIdFloat, ok := bodyUserId.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("userId is not a valid number"))
		return
	}

	// Convert float64 to int64
	userId := int64(userIdFloat)
	orderUpdateDTO.UserID = userId

	err = orderUpdateDTO.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse(err.Error()))
		return
	}

	id := c.Param("id")
	isUpdated := uc.orderService.UpdateById(id, orderUpdateDTO)
	if !isUpdated {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("Update Failed"))
		return
	}
	c.JSON(http.StatusOK, dtos.ErrorResponse("Data has been Updated"))
}

func (uc OrderController) Delete(c *gin.Context) {

	id := c.Param("id")
	isDeleted := uc.orderService.DeleteById(id)
	if !isDeleted {
		c.JSON(http.StatusBadRequest, "order cannot delete")
		return
	}
	c.JSON(http.StatusOK, "deleted")
}
