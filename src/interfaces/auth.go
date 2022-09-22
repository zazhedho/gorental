package interfaces

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

type AuthService interface {
	Login(body models.User) *helpers.Response
}
