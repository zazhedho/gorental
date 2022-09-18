package users

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

var response helpers.Response

type user_service struct {
	repo interfaces.UserRepo
}

func NewService(svc interfaces.UserRepo) *user_service {
	return &user_service{svc}
}

func (s *user_service) GetAllUsers() (*helpers.Response, error) {
	data, err := s.repo.FindAllUsers()
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *user_service) AddUser(data *models.User) (*helpers.Response, error) {
	data, err := s.repo.SaveUser(data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *user_service) UpdateUser(r *http.Request, data *models.User) (*helpers.Response, error) {
	data, err := s.repo.ChangeUser(r, data)
	if err != nil {
		res := response.ResponseJSON(404, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *user_service) DeleteUser(r *http.Request, data *models.User) (*helpers.Response, error) {
	data, err := s.repo.RemoveUser(r, data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}
