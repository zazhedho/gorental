package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/gorental/src/database/orm/models"
)

func TestGetAllUsers(t *testing.T) {
	repo := RepoUserMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Users{
		{UserId: "adf5c549-0856-4623-8609-0d67eec9f0ce", Username: "dho"},
		{UserId: "e0a4f6e5-830c-430d-824b-e55f4757926d", Username: "wew"},
	}

	repo.mock.On("FindAllUsers").Return(&dataMock, nil)
	data := service.GetAllUsers()

	result := data.Data.(*models.Users)
	for i, v := range *result {
		assert.Equal(t, dataMock[i].UserId, v.UserId, "expect id from dataMock")
	}
}
func TestGetByUsername(t *testing.T) {
	repo := RepoUserMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.User{Username: "dho"}

	var expectUsername string = "dho"
	repo.mock.On("FindByUsername", expectUsername).Return(&dataMock, nil)
	data := service.GetByUsername(dataMock.Username)

	result := data.Data.(*models.User)

	assert.Equal(t, expectUsername, result.Username, "id must be 'dho'")
}

func TestAddUser(t *testing.T) {
	repo := RepoUserMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.User{
		UserId:   "adf5c549-0856-4623-8609-0d67eec9f0cd",
		Username: "dho",
		Email:    "dho@gmail.com"}

	repo.mock.On("SaveUser", &dataMock).Return(&dataMock, nil)
	data := service.AddUser(&dataMock)

	var expectUsername string = "dho"
	result := data.Data.(*models.User)

	assert.Equal(t, expectUsername, result.Username, "id must be 'dho'")
}
