package controller

import (
	"data_person_api/helper"
	"data_person_api/model/web"
	"data_person_api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PersonControllerImpl struct {
	PersonService service.PersonService
}

func NewPersonController(personService service.PersonService) PersonController {
	return &PersonControllerImpl{
		PersonService: personService,
	}
}

func (controller *PersonControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	personCreateRequest := web.PersonCreateRequest{}
	helper.ReadRequestBody(request, &personCreateRequest)

	personResponse := controller.PersonService.Create(request.Context(), personCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   personResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PersonControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	personUpdateRequest := web.PersonUpdateRequest{}
	helper.ReadRequestBody(request, &personUpdateRequest)

	personId := params.ByName("personId")
	id, err := strconv.Atoi(personId)
	helper.PanicError(err)

	personUpdateRequest.Id = id

	personResponse := controller.PersonService.Update(request.Context(), personUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   personResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PersonControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	personId := params.ByName("personId")
	id, err := strconv.Atoi(personId)
	helper.PanicError(err)

	controller.PersonService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PersonControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	personId := params.ByName("personId")
	id, err := strconv.Atoi(personId)
	helper.PanicError(err)

	personResponse := controller.PersonService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   personResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PersonControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	personResponses := controller.PersonService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   personResponses,
	}

	helper.WriteResponseBody(writer, webResponse)
}
