package users

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/interfaces"
)

type user_service struct {
	repo interfaces.UserRepo
}

func NewService(repo interfaces.UserRepo) *user_service {
	return &user_service{repo}
}

func (s *user_service) GetAllUsers() (*models.Users, error) {
	data, err := s.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *user_service) AddUser(data *models.User) (*models.User, error) {
	data, err := s.repo.SaveUser(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *user_service) UpdateUser(r *http.Request, data *models.User) (*models.User, error) {
	data, err := s.repo.ChangeUser(r, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *user_service) DeleteUser(r *http.Request, data *models.User) (*models.User, error) {
	data, err := s.repo.RemoveUser(r, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
