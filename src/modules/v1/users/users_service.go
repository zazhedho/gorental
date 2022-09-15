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

func (r *user_service) GetAllUsers() (*models.Users, error) {
	data, err := r.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *user_service) AddUser(data *models.User) (*models.User, error) {
	data, err := r.repo.SaveUser(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *user_service) UpdateUser(re *http.Request, data *models.User) (*models.User, error) {
	data, err := r.repo.ChangeUser(re, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *user_service) DeleteUser(re *http.Request, data *models.User) (*models.User, error) {
	data, err := r.repo.RemoveUser(re, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
