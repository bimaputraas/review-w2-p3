package repository

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBRepository struct {
	DB *mongo.Database
}

func NewMongoDBRepository(db *mongo.Database) Repository {
	return MongoDBRepository{DB: db}
}