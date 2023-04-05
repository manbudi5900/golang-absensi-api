package main

import (
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

	userService.SaveAvatar("cd3c5838-5acb-43b9-bb21-de5ee511a3d1", "images/cd3c5838-5acb-43b9-bb21-de5ee511a3d1-profile.png")

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()
}
