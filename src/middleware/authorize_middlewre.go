package middleware

import (
	"api-customer/dtos"
	"api-customer/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fmt.Println(authorizationHeader)
		if authorizationHeader == "" {
			fmt.Println(dtos.ErrorResponse("Token in header cannot be blank"))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrorResponse("Token in header cannot be blank"))
			return
		}

		if !strings.Contains(authorizationHeader, "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, dtos.ErrorResponse("Token format is invalid"))
			return
		}

		claims, err := services.ValidateToken(authorizationHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, dtos.ErrorResponse(err.Error()))
			return
		}

		ctx.Set("userId", claims["id"])
		ctx.Set("userInfo", claims)
		ctx.Next()
	}
}
