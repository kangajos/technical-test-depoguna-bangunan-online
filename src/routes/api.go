package routes

import (
	"api-customer/controllers"
	"api-customer/middleware"
	"api-customer/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// health := new(controllers.HealthController)

	// router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			auth := controllers.NewAuthController(*services.NewAuthService(db))
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/logout", auth.Logout)
		}

		v1.Use(middleware.AuthorizeMiddleware())
		userGroup := v1.Group("user")
		{
			user := controllers.NewUserController(*services.NewUserService(db))
			userGroup.GET("/", user.Pagination)
			userGroup.GET("/:id", user.FindById)
			userGroup.POST("/", user.Create)
			userGroup.PUT("/:id", user.Update)
			userGroup.DELETE("/:id", user.Delete)
		}

		orderGroup := v1.Group("order")
		{
			order := controllers.NewOrderController(*services.NewOrderService(db))
			orderGroup.GET("/", order.Pagination)
			orderGroup.GET("/:id", order.FindById)
			orderGroup.POST("/", order.Create)
			orderGroup.PUT("/:id", order.Update)
			orderGroup.DELETE("/:id", order.Delete)
		}
	}
	return router

}
