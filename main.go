package main

import (
	"data_person_api/app"
	"data_person_api/controller"
	"data_person_api/helper"
	"data_person_api/middleware"
	"data_person_api/repository"
	"data_person_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	personRepository := repository.NewPersonRepository()
	personService := service.NewPersonService(personRepository, db, validate)
	personController := controller.NewPersonController(personService)
	router := app.NewRouter(personController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
