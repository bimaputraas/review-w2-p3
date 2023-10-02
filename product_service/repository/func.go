package repository

import (
	"context"
	"preview-w2/product_service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongodb) CreateProduct(product model.Product) (model.Product, error) {
	// get coll
	result,err := r.db.Collection("products").InsertOne(context.Background(),product)
	if err != nil {
		return model.Product{},err
	}

	// insert id to model
	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}

func (r *mongodb) FindProducts() ([]model.Product, error){
	return nil,nil
}

func (r *mongodb) FindProductById(int) (model.Product, error){
	return model.Product{},nil
}

func (r *mongodb) UpdateProductById(int, model.Product) (model.Product, error){
	return model.Product{},nil
}

func (r *mongodb) DeleteProductById(int) (model.Product, error){
	return model.Product{},nil
}

