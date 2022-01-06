package app

import (
	"data_person_api/controller"
	"data_person_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(personController controller.PersonController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/persons", personController.FindAll)
	router.GET("/persons/:personId", personController.FindById)
	router.POST("/persons", personController.Create)
	router.PUT("/persons/:personId", personController.Update)
	router.DELETE("/persons/:personId", personController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
