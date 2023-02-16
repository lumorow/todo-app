package main

import (
	"github.com/lumorow/todo-app"
	"github.com/lumorow/todo-app/pkg/handler"
	"github.com/lumorow/todo-app/pkg/repository"
	"github.com/lumorow/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
