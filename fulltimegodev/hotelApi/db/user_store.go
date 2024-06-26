package db

import (
	"github.com/alikazai/golearn/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserById(string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
	}
}

func (m *MongoUserStore) GetUserById(id string) (*types.User, error) {
	return nil, nil
}
