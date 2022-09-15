package interfaces

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
)

type UserRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(re *http.Request, data *models.User) (*models.User, error)
	RemoveUser(re *http.Request, data *models.User) (*models.User, error)
}

type UserServive interface {
	GetAllUsers() (*models.Users, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(re *http.Request, data *models.User) (*models.User, error)
	DeleteUser(re *http.Request, data *models.User) (*models.User, error)
}
