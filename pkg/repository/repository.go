package repository

import (
	"github.com/jmoiron/sqlx"
	"shop"
)

type Authorization interface {
	CreateUser(user shop.User) (int, error)
	GetUser(login, password string) (shop.User, error)
}

type Product interface {
	CreateProduct(product shop.Product) (int, error)
	GetAllProducts() ([]shop.Product, error)
	GetProductById(id int) (shop.Product, error)
}

type Repository struct {
	Authorization
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Product:       NewProductListPostgres(db),
	}
}
