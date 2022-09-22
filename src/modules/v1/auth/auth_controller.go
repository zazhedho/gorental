package auth

import (
	"encoding/json"
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type auth_ctrl struct {
	repo interfaces.AuthService
}

func NewCtrl(ctrl interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{ctrl}
}

func (a auth_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err.Error(), 401, true)
		return
	}

	a.repo.Login(data).Send(w)
}
