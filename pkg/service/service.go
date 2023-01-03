package service

import (
	"shop"
	"shop/pkg/repository"
)

type Authorization interface {
	CreateUser(user shop.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Product interface {
	CreateProduct(product shop.Product) (int, error)
	GetAllProducts() ([]shop.Product, error)
	GetProductById(id int) (shop.Product, error)
}

type Service struct {
	Authorization
	Product
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
		Product:       NewProductListService(r.Product),
	}
}
