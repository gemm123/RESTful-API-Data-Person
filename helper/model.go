package helper

import (
	"data_person_api/model/domain"
	"data_person_api/model/web"
)

func ToPersonResponse(person domain.Person) web.PersonResponse {
	return web.PersonResponse{
		Id:        person.Id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Age:       person.Age,
	}
}

func ToPersonResponses(persons []domain.Person) []web.PersonResponse {
	var PersonReponses []web.PersonResponse
	for _, person := range persons {
		PersonReponses = append(PersonReponses, ToPersonResponse(person))
	}

	return PersonReponses
}
