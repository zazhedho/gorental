package interfaces

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

type UserRepo interface {
	FindAllUsers() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	UserExists(username, email string) bool
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(username string, data *models.User) (*models.User, error)
	RemoveUser(username string, data *models.User) (*models.User, error)
}

type UserService interface {
	GetAllUsers() *helpers.Response
	GetByUsername(username string) *helpers.Response
	AddUser(data *models.User) *helpers.Response
	UpdateUser(username string, data *models.User) *helpers.Response
	DeleteUser(username string, data *models.User) *helpers.Response
}
