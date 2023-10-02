package repository

import "preview-w2/product_service/model"

type Repository interface {
	CreateProduct(model.Product) (model.Product, error)
	FindProducts() ([]model.Product, error)
	FindProductById(int) (model.Product, error)
	UpdateProductById(int, model.Product) (model.Product, error)
	DeleteProductById(int) (model.Product, error)
}