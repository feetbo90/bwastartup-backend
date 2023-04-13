package main

import (
	"bwastartup/user"
	"log"

	"bwastartup/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	// user := user.User{
	// 	Name: "Test Simpan",
	// }

	userService := user.NewService(userRepository)

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test dari service"

	// userService.RegisterUser(userInput)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()
}
