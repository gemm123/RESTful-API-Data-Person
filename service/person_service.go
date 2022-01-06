package service

import (
	"context"
	"data_person_api/model/web"
)

type PersonService interface {
	Create(ctx context.Context, request web.PersonCreateRequest) web.PersonResponse
	Update(ctx context.Context, request web.PersonUpdateRequest) web.PersonResponse
	Delete(ctx context.Context, personId int)
	FindById(ctx context.Context, personId int) web.PersonResponse
	FindAll(ctx context.Context) []web.PersonResponse
}
