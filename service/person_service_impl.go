package service

import (
	"context"
	"data_person_api/exception"
	"data_person_api/helper"
	"data_person_api/model/domain"
	"data_person_api/model/web"
	"data_person_api/repository"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type PersonServiceImpl struct {
	PersonRepository repository.PersonRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewPersonService(personRepository repository.PersonRepository, DB *sql.DB, validate *validator.Validate) PersonService {
	return &PersonServiceImpl{
		PersonRepository: personRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *PersonServiceImpl) Create(ctx context.Context, request web.PersonCreateRequest) web.PersonResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)

	person := domain.Person{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Age:       request.Age,
	}

	person = service.PersonRepository.Save(ctx, tx, person)

	return helper.ToPersonResponse(person)
}

func (service *PersonServiceImpl) Update(ctx context.Context, request web.PersonUpdateRequest) web.PersonResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)

	person, err := service.PersonRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	person.FirstName = request.FirstName
	person.LastName = request.LastName
	person.Age = request.Age

	person = service.PersonRepository.Update(ctx, tx, person)

	return helper.ToPersonResponse(person)
}

func (service *PersonServiceImpl) Delete(ctx context.Context, personId int) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)

	person, err := service.PersonRepository.FindById(ctx, tx, personId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PersonRepository.Delete(ctx, tx, person)
}

func (service *PersonServiceImpl) FindById(ctx context.Context, personId int) web.PersonResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)

	person, err := service.PersonRepository.FindById(ctx, tx, personId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPersonResponse(person)
}

func (service *PersonServiceImpl) FindAll(ctx context.Context) []web.PersonResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(tx)

	persons := service.PersonRepository.FindAll(ctx, tx)

	return helper.ToPersonResponses(persons)
}
