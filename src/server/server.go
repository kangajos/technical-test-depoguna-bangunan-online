package server

import (
	"api-customer/models"
	"api-customer/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	db  *gorm.DB
	gin gin.Engine
}

func (s *Server) DBConnection(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	s.db = db
	s.DBMigration()
}

func (s *Server) DBMigration() {
	s.db.AutoMigrate(models.User{}, models.Order{})
	s.db.AutoMigrate(models.User{}, models.Order{})
}

func (s *Server) Run() {
	r := routes.NewRouter(s.db)
	r.Run()
}
