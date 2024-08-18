package main

import (
	"api-customer/server"
	"fmt"
)

func main() {
	fmt.Println("welcome to golang api")

	dsn := "admin:Querty123!@tcp(127.0.0.1:3306)/tech_test_dbo?charset=utf8mb4&parseTime=True&loc=Local"

	var srv server.Server
	srv.DBConnection(dsn)
	srv.Run()
}
