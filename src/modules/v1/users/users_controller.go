package users

import (
	"encoding/json"
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/interfaces"
)

type user_ctrl struct {
	srv interfaces.UserServive
}

func NewCtrl(repo interfaces.UserServive) *user_ctrl {
	return &user_ctrl{srv: repo}
}

func (re *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := re.srv.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (re *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := re.srv.AddUser(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (re *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := re.srv.UpdateUser(r, &datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (re *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User

	data, err := re.srv.DeleteUser(r, &datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}
