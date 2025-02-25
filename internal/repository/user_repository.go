package repository

import (
	"GO-CRUD-PROJECT/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "log"
	"time"
)

func CreateUser(user model.User) (model.User, error) {
	collection := GetMongoCollection()
	user.ID = int(time.Now().Unix()) // Generate a unique ID (you can use a better approach here)
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUser(id int) (model.User, error) {
	collection := GetMongoCollection()
	var user model.User
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return model.User{}, nil // Not found
	}
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]model.User, error) {
	collection := GetMongoCollection()
	var users []model.User
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var user model.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func UpdateUser(id int, user model.User) (model.User, error) {
	collection := GetMongoCollection()
	result := collection.FindOneAndUpdate(
		context.Background(),
		bson.M{"id": id},
		bson.M{
			"$set": bson.M{
				"name":  user.Name,
				"email": user.Email,
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)
	var updatedUser model.User
	err := result.Decode(&updatedUser)
	if err != nil {
		return model.User{}, err
	}
	return updatedUser, nil
}

func DeleteUser(id int) error {
	collection := GetMongoCollection()
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": id})
	return err
}
