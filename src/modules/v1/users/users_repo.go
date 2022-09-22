package users

import (
	"errors"

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

func (re *user_repo) FindByUsername(username string) (*models.User, error) {
	var data models.User
	result := re.db.First(&data, "username = ?", username)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (re *user_repo) UserExists(username, email string) bool {
	var data models.User
	result := re.db.First(&data, "username = ? OR email = ?", username, email)
	return result.Error == nil
}

func (re *user_repo) SaveUser(data *models.User) (*models.User, error) {
	result := re.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal menambah data")
	}

	return data, nil
}

func (re *user_repo) ChangeUser(username string, data *models.User) (*models.User, error) {

	result := re.db.Model(&data).Where("username = ?", username).Updates(data)
	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data user tidak ditemukan")
	}

	return data, nil
}

func (re *user_repo) RemoveUser(username string, data *models.User) (*models.User, error) {
	result := re.db.Where("username = ?", username).Delete(data)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data user tidak ditemukan")
	}

	return nil, nil
}
