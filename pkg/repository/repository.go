package repository

import "github.com/jmoiron/sqlx"

type Email interface {
	Create(email string) error
	Delete(email string) error
	GetAll() ([]string, error)
}

type Repository struct {
	Email
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Email: NewEmailPostgres(db),
	}
}
