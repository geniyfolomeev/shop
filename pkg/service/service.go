package service

import "shop/pkg/repository"

type Authorization interface {
}

type Product interface {
}

type Service struct {
	Authorization
	Product
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
