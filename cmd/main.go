package main

import (
	taxi "TaxiApp1"
	"TaxiApp1/config"
	"TaxiApp1/internal/handler"
	"TaxiApp1/internal/repository"
	"TaxiApp1/internal/service"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//load config with environment variables
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println(config.DBDriver)

	//db connection
	db, err := repository.NewPostgresDB(repository.ConfigDB{
		DBDriver: config.DBDriver,
		DBSource: config.DBSource,
	})
	if err != nil {
		log.Fatalf("failed to initialize db:%s", err.Error())
	}
	//repository
	repos := repository.NewRepository(db)
	//services
	services := service.NewService(repos)
	//handlers
	handlers := handler.NewHandler(services)
	//server
	srv := new(taxi.Server)
	if err := srv.Run(config.ServerPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}

}
