package users

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type user_service struct {
	repo interfaces.UserRepo
}

func NewService(svc interfaces.UserRepo) *user_service {
	return &user_service{svc}
}

func (s *user_service) GetAllUsers() *helpers.Response {
	result, err := s.repo.FindAllUsers()
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 200, false)
}

func (s *user_service) GetByUsername(username string) *helpers.Response {
	result, err := s.repo.FindByUsername(username)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}
	return helpers.New(result, 200, false)
}

func (s *user_service) AddUser(data *models.User) *helpers.Response {
	if check := s.repo.UserExists(data.Username, data.Email); check {
		return helpers.New("username or email already exists", 400, true)
	}

	hashPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	data.Password = hashPassword
	result, err := s.repo.SaveUser(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	return helpers.New(result, 201, false)
}

func (s *user_service) UpdateUser(username string, data *models.User) *helpers.Response {
	if check := s.repo.UserExists(data.Username, data.Email); check {
		return helpers.New("username sudah terdaftar", 400, true)
	}

	hashPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	data.Password = hashPassword
	result, err := s.repo.ChangeUser(username, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}
	return helpers.New(result, 200, false)
}

func (s *user_service) DeleteUser(username string, data *models.User) *helpers.Response {
	result, err := s.repo.RemoveUser(username, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}
	return helpers.New(result, 200, false)
}
