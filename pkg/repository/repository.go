package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Product interface {
}

type Repository struct {
	Authorization
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
