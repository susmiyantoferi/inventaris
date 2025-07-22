package main

import (
	"inventaris/app"
	"inventaris/controller"
	"inventaris/repository"
	"inventaris/route"
	"inventaris/service"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

func main(){
	err := os.MkdirAll("uploads/", os.ModePerm)
	if err != nil {
		log.Fatal("Gagal membuat folder uploads", err)
	}

	database := app.Db()
	validate := validator.New()
	
	produkRepo := repository.NewProdukRepositoryImpl(database)
	produkService := service.NewProdukServiceImpl(produkRepo, validate)
	produkController := controller.NewProdukControllerImpl(produkService)

	inventRepo := repository.NewInventarisRepositoryImpl(database)
	inventService := service.NewInventarisServImpl(inventRepo, validate, database)
	inventController := controller.NewInventControllerImpl(inventService)

	pesananRepo := repository.NewPesananRepositoryImpl(database)
	pesananService := service.NewPesananServiceImpl(pesananRepo, validate, produkRepo)
	pesananController := controller.NewPesananControllerImpl(pesananService)

	routes := route.NewRouter(produkController, inventController, pesananController)

	//testtt triger ci/cd
	
	log.Println("Server run at http://localhost:8080")
	routes.Run(":8080")
	
}