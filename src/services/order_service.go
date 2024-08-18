package services

import (
	"api-customer/dtos"
	"api-customer/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (us OrderService) Pagination(c *gin.Context, userId int64) paginate.Page {
	var order []models.Order
	stmt := us.db.Model(order).Where("user_id = ?", userId)
	if c.Query("name") != "" {
		stmt.Where("name like ?", "%"+c.Query("name")+"%")
	}
	pg := paginate.New()
	paginate := pg.With(stmt).Request(c.Request).Response(&order)
	return paginate
}

func (us OrderService) FindById(id string) (order models.Order, err error) {
	ID, _ := strconv.ParseInt(id, 10, 64)
	err = us.db.Model(order).First(&order, "id = ?", ID).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("get order by id %d err: %s", id, err.Error()))
		return order, err
	}
	return order, nil
}

func (us OrderService) Create(body dtos.OrderCreateDTO) (order models.Order, err error) {
	order.Name = body.Name
	order.Qty = body.Qty
	order.Price = body.Price
	order.UserID = body.UserID
	err = us.db.Save(&order).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("create order err: %s", err.Error()))
		return order, err
	}
	return order, nil
}

func (us OrderService) UpdateById(id string, body dtos.OrderUpdateDTO) bool {
	order, err := us.FindById(id)
	if err != nil {
		return false
	}

	if order.UserID != body.UserID {
		return false
	}

	order.Name = body.Name
	order.Qty = body.Qty
	order.Price = body.Price
	order.UserID = body.UserID
	err = us.db.Save(&order).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("update order err: %s", err.Error()))
		return false
	}

	return true
}

func (us OrderService) DeleteById(id string) bool {
	ID, _ := strconv.ParseInt(id, 10, 64)
	err := us.db.Delete(&models.Order{}, ID).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("delete order by id %d err: %s", id, err.Error()))
		return false
	}

	return true
}
