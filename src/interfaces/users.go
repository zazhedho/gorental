package interfaces

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
)

type UserRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(r *http.Request, data *models.User) (*models.User, error)
	RemoveUser(r *http.Request, data *models.User) (*models.User, error)
}

type UserService interface {
	GetAllUsers() (*models.Users, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(r *http.Request, data *models.User) (*models.User, error)
	DeleteUser(r *http.Request, data *models.User) (*models.User, error)
}
