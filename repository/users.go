package repository

import (
	"context"

	"github.com/scratchingmycranium/golang-rest/interfaces"
	"github.com/scratchingmycranium/golang-rest/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func InitUserRepo(client *mongo.Client, db string, col string) interfaces.IUserRepo {
	collection := client.Database(db).Collection(col)

	return &UserRepo{
		Collection: collection,
	}
}

func (u *UserRepo) Get() (model.User, error) {
	// Create a filter to find the document you want to get
	filter := bson.D{{"name", "Ash"}}

	// Create a value into which the result can be decoded
	var result model.User

	err := u.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil

}
