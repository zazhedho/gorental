package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(ctrl interfaces.UserService) *user_ctrl {
	return &user_ctrl{ctrl}
}

func (c *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllUsers().Send(w)
}
func (c *user_ctrl) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	c.svc.GetByUsername(username).Send(w)
}

func (c *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {

	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}
	c.svc.AddUser(&data).Send(w)
}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	username := r.Context().Value("username").(string)
	c.svc.UpdateUser(username, &data).Send(w)
}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var data models.User

	params := mux.Vars(r)
	c.svc.DeleteUser(params["username"], &data).Send(w)
}
