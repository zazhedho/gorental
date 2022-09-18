package users

import (
	"encoding/json"
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/interfaces"
)

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
		response.ResponseJSON(400, "Tidak dapat menampilkan data").Send(w)
	}

	json.NewEncoder(w).Encode(data)
}

func (c *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := c.svc.AddUser(&datas)
	if err != nil {
		response.ResponseJSON(400, "Tidak dapat menambahkan data").Send(w)
	}

	json.NewEncoder(w).Encode(data)
}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := c.svc.UpdateUser(r, &datas)
	if err != nil {
		response.ResponseJSON(400, "Tidak dapat memperbarui data").Send(w)
	}

	json.NewEncoder(w).Encode(data)
}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User

	data, err := c.svc.DeleteUser(r, &datas)
	if err != nil {
		response.ResponseJSON(400, "Tidak dapat menghapus data").Send(w)
	}

	json.NewEncoder(w).Encode(data)
}
