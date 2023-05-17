package repository

type Authorization interface {
}

type Product interface {
}

type Repository struct {
	Authorization
	Product
}

func NewRepository() *Repository {
	return &Repository{}
}
