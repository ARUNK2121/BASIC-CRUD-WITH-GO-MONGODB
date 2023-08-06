package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-mongo-sample/helper"
	"go-mongo-sample/model"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r repository) GetUser(ctx context.Context, email string) (model.UserJSON, error) {
	var out model.UserBSON
	err := r.db.
		Collection("users").
		FindOne(ctx, bson.M{"email": email}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.UserJSON{}, ErrUserNotFound
		}
		return model.UserJSON{}, err
	}
	return helper.BsonToJson(out), nil
}

func (r repository) CreateUser(ctx context.Context, user model.UserJSON) (model.UserJSON, error) {
	out, err := r.db.
		Collection("users").
		InsertOne(ctx, helper.JsonToBson(user))
	if err != nil {
		return model.UserJSON{}, err
	}
	user.ID = out.InsertedID.(primitive.ObjectID).String()
	return user, nil
}

func (r repository) UpdateUser(ctx context.Context, user model.UserJSON) (model.UserJSON, error) {
	in := bson.M{}
	if user.Name != "" {
		in["name"] = user.Name
	}
	if user.Password != "" {
		in["password"] = user.Password
	}
	out, err := r.db.
		Collection("users").
		UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": in})
	if err != nil {
		return model.UserJSON{}, err
	}
	if out.MatchedCount == 0 {
		return model.UserJSON{}, ErrUserNotFound
	}
	return user, nil
}

func (r repository) DeleteUser(ctx context.Context, email string) error {
	out, err := r.db.
		Collection("users").
		DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}
	if out.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}
