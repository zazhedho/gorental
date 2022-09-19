package users

import (
	"encoding/json"
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

var response helpers.Response

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(ctrl interfaces.UserService) *user_ctrl {
	return &user_ctrl{ctrl}
}

func (c *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := c.svc.GetAllUsers()
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menampilkan data", data, err)
	} else {
		response.ResponseJSON(w, 200, "successfully", data, err)
	}
}

func (c *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		data, err := c.svc.AddUser(&datas)
		if err != nil {
			response.ResponseJSON(w, 400, "", data, err)
		} else {

			response.ResponseJSON(w, 201, "user has been created", data, err)
		}
	}
}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		data, err := c.svc.UpdateUser(r, &datas)
		if err != nil {
			response.ResponseJSON(w, 400, "Tidak dapat memperbarui data", data, err)
			return
		} else {

			response.ResponseJSON(w, 200, "user has been updated", data, err)
		}
	}
}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	data, err := c.svc.DeleteUser(r, &datas)
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menghapus data", data, err)
	} else {
		response.ResponseJSON(w, 200, "deleted successfully", data, err)
	}
}
