package main

import (
	"Makves/handler"
	"Makves/repository"
	"Makves/usecase"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	"log"

	_ "Makves/swagger"
)

func initConfig() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @host 127.0.0.1:8000
func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	userRepo, err := repository.NewSCVRepo(viper.GetString("filePath"))
	if err != nil {
		log.Fatal(err)
	}
	uc := usecase.NewUserUseCase(userRepo)
	handler := handler.NewHandler(uc)

	app := fiber.New()

	swaggerConf := swagger.ConfigDefault
	app.Get("/swagger/*", swagger.New(swaggerConf))

	app.Get("/get-items", handler.GetUserByIds)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", viper.GetInt("port"))))
}
