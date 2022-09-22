package auth

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type auth_service struct {
	repo interfaces.UserRepo
}

type token_response struct {
	Tokens string `json:"token"`
}

func NewService(svc interfaces.UserRepo) *auth_service {
	return &auth_service{svc}
}

func (a auth_service) Login(body models.User) *helpers.Response {
	user, err := a.repo.FindByUsername(body.Username)
	if err != nil {
		return helpers.New("username tidak terdaftar", 401, true)
	}

	if !helpers.CheckPassword(user.Password, body.Password) {
		return helpers.New("password is wrong", 401, true)
	}

	token := helpers.NewToken(body.Username, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return helpers.New(err.Error(), 401, true)
	}
	return helpers.New(token_response{Tokens: theToken}, 200, false)
}
