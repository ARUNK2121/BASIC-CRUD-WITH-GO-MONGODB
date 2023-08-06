package helper

import (
	"go-mongo-sample/model"
)

func JsonToBson(in model.UserJSON) model.UserBSON {
	return model.UserBSON{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}

func BsonToJson(in model.UserBSON) model.UserJSON {
	return model.UserJSON{
		ID:       in.ID.String(),
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}
