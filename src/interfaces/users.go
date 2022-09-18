package interfaces

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

type UserRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(r *http.Request, data *models.User) (*models.User, error)
	RemoveUser(r *http.Request, data *models.User) (*models.User, error)
}

type UserService interface {
	GetAllUsers() (*helpers.Response, error)
	AddUser(data *models.User) (*helpers.Response, error)
	UpdateUser(r *http.Request, data *models.User) (*helpers.Response, error)
	DeleteUser(r *http.Request, data *models.User) (*helpers.Response, error)
}
