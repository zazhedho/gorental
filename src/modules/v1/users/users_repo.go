package users

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/database/orm/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (r *user_repo) FindAllUsers() (*models.Users, error) {
	var data models.Users

	result := r.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *user_repo) SaveUser(data *models.User) (*models.User, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambah data")
	}

	return data, nil
}

func (r *user_repo) ChangeUser(re *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(re)
	result := r.db.Model(&data).Where("name = ?", vars["name"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	return data, nil
}

func (r *user_repo) RemoveUser(re *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(re)
	result := r.db.Where("name = ?", vars["name"]).Delete(&data)

	if result.Error != nil {
		return nil, errors.New("gagal hapus data")
	}

	return data, nil
}
