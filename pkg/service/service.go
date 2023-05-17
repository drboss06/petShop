package service

import "petShop/pkg/repository"

type Authorization interface {
}

type Product interface {
}

type Service struct {
	Authorization
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
