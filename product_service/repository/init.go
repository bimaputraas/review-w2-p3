package repository

import "go.mongodb.org/mongo-driver/mongo"

// mongodb repository
type mongodb struct {
	db *mongo.Database
}

func NewMongoDbRepository(Db *mongo.Database) Repository {
	return &mongodb{db: Db}
}