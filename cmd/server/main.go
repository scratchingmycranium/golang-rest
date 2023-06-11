package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/scratchingmycranium/golang-rest/config"
	"github.com/scratchingmycranium/golang-rest/http"
	"github.com/scratchingmycranium/golang-rest/repository"
	"github.com/scratchingmycranium/golang-rest/service"
	"github.com/scratchingmycranium/golang-rest/utils"
)

func main() {

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config := config.InitConfig()

	utils := utils.InitUtils(config)

	mongoClient := utils.InitMongo()

	router := utils.InitRouter()

	userRepo := repository.InitUserRepo(mongoClient, config.MongoDB, config.MongoUserCollection)
	userSvc := service.InitUserService(userRepo)
	userHandler := http.InitUserHandler(userSvc)

	userHandler.RegisterUserRoutes(router)

	router.Run()
}
