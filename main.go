package main

import (
	"inventaris/app"
	"inventaris/controller"
	"inventaris/repository"
	"inventaris/route"
	"inventaris/service"
	"log"

	"github.com/go-playground/validator/v10"
)

func main(){

	database := app.Db()
	validate := validator.New()
	
	produkRepo := repository.NewProdukRepositoryImpl(database)
	produkService := service.NewProdukServiceImpl(produkRepo, validate)
	produkController := controller.NewProdukControllerImpl(produkService)

	inventRepo := repository.NewInventarisRepositoryImpl(database)
	inventService := service.NewInventarisServImpl(inventRepo, validate, database)
	inventController := controller.NewInventControllerImpl(inventService)

	routes := route.NewRouter(produkController, inventController)
	
	log.Println("Server run at http://localhost:8080")
	routes.Run(":8080")
	
}