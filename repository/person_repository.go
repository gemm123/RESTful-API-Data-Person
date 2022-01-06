package repository

import (
	"context"
	"data_person_api/model/domain"
	"database/sql"
)

type PersonRepository interface {
	Save(ctx context.Context, tx *sql.Tx, person domain.Person) domain.Person
	Update(ctx context.Context, tx *sql.Tx, person domain.Person) domain.Person
	Delete(ctx context.Context, tx *sql.Tx, person domain.Person)
	FindById(ctx context.Context, tx *sql.Tx, personId int) (domain.Person, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Person
}
