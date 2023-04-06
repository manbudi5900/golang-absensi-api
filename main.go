package main

import (
	"absensi/auth"
	"absensi/handler"
	"absensi/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dbURL := "postgres://postgres:root@localhost:5432/absensi"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connection to database is good")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()
}
