package app

import (
	"fmt"
	"os"
	"sesi_8/config"
	"sesi_8/handler"
	"sesi_8/repository"
	"sesi_8/route"
	"sesi_8/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
