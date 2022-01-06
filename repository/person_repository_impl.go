package repository

import (
	"context"
	"data_person_api/helper"
	"data_person_api/model/domain"
	"database/sql"
	"errors"
)

type PersonRepositoryImpl struct {
}

func NewPersonRepository() PersonRepository {
	return &PersonRepositoryImpl{}
}

func (repository *PersonRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, person domain.Person) domain.Person {
	SQL := "INSERT INTO person(first_name, last_name, age) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, person.FirstName, person.LastName, person.Age)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	person.Id = int(id)
	return person
}

func (repository *PersonRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, person domain.Person) domain.Person {
	SQL := "UPDATE person SET first_name = ?, last_name = ?, age = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, person.FirstName, person.LastName, person.Age, person.Id)
	helper.PanicError(err)

	return person
}

func (repository *PersonRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, person domain.Person) {
	SQL := "DELETE FROM person WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, person.Id)
	helper.PanicError(err)
}

func (repository *PersonRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, personId int) (domain.Person, error) {
	SQL := "SELECT id, first_name, last_name, age FROM person WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, personId)
	helper.PanicError(err)
	defer rows.Close()

	person := domain.Person{}
	if rows.Next() {
		err := rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age)
		helper.PanicError(err)
		return person, nil
	} else {
		return person, errors.New("person is not found")
	}
}

func (repository *PersonRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Person {
	SQL := "SELECT id, first_name, last_name, age FROM person"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var persons []domain.Person
	for rows.Next() {
		person := domain.Person{}
		err := rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age)
		helper.PanicError(err)
		persons = append(persons, person)
	}
	return persons
}
