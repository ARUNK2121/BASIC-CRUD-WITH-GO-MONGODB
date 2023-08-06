package repository

import (
	"context"

	"go-mongo-sample/model"
)

type Repository interface {
	GetUser(ctx context.Context, email string) (model.UserJSON, error)
	CreateUser(ctx context.Context, in model.UserJSON) (model.UserJSON, error)
	UpdateUser(ctx context.Context, in model.UserJSON) (model.UserJSON, error)
	DeleteUser(ctx context.Context, email string) error
}
