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

func (re *user_repo) FindAllUsers() (*models.Users, error) {
	var data models.Users

	result := re.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (re *user_repo) SaveUser(data *models.User) (*models.User, error) {
	result := re.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambah data")
	}

	return data, nil
}

func (re *user_repo) ChangeUser(r *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(r)
	result := re.db.Model(&data).Where("name = ?", vars["name"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	return data, nil
}

func (re *user_repo) RemoveUser(r *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(r)
	result := re.db.Where("name = ?", vars["name"]).Delete(&data)

	if result.Error != nil {
		return nil, errors.New("gagal hapus data")
	}

	return data, nil
}
