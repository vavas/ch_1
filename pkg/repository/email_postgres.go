package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EmailPostgres struct {
	db *sqlx.DB
}

func NewEmailPostgres(db *sqlx.DB) *EmailPostgres {
	return &EmailPostgres{db: db}
}

func (r *EmailPostgres) Create(email string) error {
	createEmailQuery := fmt.Sprintf("INSERT INTO %s (email) values ($1)", emailsTable)
	_, err := r.db.Exec(createEmailQuery, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *EmailPostgres) GetAll() ([]string, error) {
	var emails []string
	query := fmt.Sprintf(`SELECT email FROM %s`, emailsTable)
	if err := r.db.Select(&emails, query); err != nil {
		return nil, err
	}
	return emails, nil
}

func (r *EmailPostgres) Delete(email string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE email = $1`, emailsTable)
	_, err := r.db.Exec(query, email)
	return err
}
