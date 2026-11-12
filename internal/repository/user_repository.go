package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (r *UserRepository) CreateUser(user interface{}) error {
	_, err := r.Collection.InsertOne(context.Background(), user)
	return err
}
