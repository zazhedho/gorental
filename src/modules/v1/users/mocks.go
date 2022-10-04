package users

import (
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/gorental/src/database/orm/models"
)

type RepoUserMock struct {
	mock mock.Mock
}

func (m *RepoUserMock) FindAllUsers() (*models.Users, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Users), nil
}

func (m *RepoUserMock) FindByUsername(username string) (*models.User, error) {
	args := m.mock.Called(username)
	return args.Get(0).(*models.User), nil
}

func (m *RepoUserMock) UserExists(username, email string) bool {
	return false
}

func (m *RepoUserMock) SaveUser(data *models.User) (*models.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.User), nil
}

func (m *RepoUserMock) ChangeUser(username string, data *models.User) (*models.User, error) {
	args := m.mock.Called(username, data)
	return args.Get(0).(*models.User), nil
}

func (m *RepoUserMock) RemoveUser(username string, data *models.User) (*models.User, error) {
	args := m.mock.Called(username, data)
	return args.Get(0).(*models.User), nil
}
