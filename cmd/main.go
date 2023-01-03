package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"shop"
	"shop/pkg/handler"
	"shop/pkg/repository"
	"shop/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config init error: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		log.Fatalf("DB init failed: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(shop.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Critical error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
