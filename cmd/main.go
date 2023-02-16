package main

import (
	_ "github.com/lib/pq"
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

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigFile("config")
	//viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
