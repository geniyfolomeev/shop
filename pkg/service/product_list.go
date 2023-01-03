package service

import (
	"shop"
	"shop/pkg/repository"
)

type ProductListService struct {
	repo repository.Product
}

func NewProductListService(repo repository.Product) *ProductListService {
	return &ProductListService{repo: repo}
}

func (s *ProductListService) CreateProduct(product shop.Product) (int, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductListService) GetAllProducts() ([]shop.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductListService) GetProductById(id int) (shop.Product, error) {
	return s.repo.GetProductById(id)
}
