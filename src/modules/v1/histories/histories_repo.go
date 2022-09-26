package histories

import (
	"errors"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (re *history_repo) FindAllHistories() (*models.Histories, error) {
	var data models.Histories

	result := re.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return re.db.Select("user_id", "username", "created_at", "updated_at")
	}).Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
		return re.db.Select("vehicle_id", "name", "rating", "created_at", "updated_at")
	}).Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &data, nil
}

func (re *history_repo) SaveHistory(data *models.History) (*models.History, error) {
	result := re.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal menambah data")
	}

	return data, nil
}

func (re *history_repo) ChangeHistory(id int, data *models.History) (*models.History, error) {

	result := re.db.Model(&data).Where("history_id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("id history tidak ditemukan")
	}

	return data, nil
}

func (re *history_repo) RemoveHistory(id int, data *models.History) (*models.History, error) {
	result := re.db.Delete(data, id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("id history tidak ditemukan")
	}

	return nil, nil
}

// Search
func (re *history_repo) FindHistoryByVehicleId(id int, data *models.Histories) (*models.Histories, error) {

	result := re.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return re.db.Select("user_id", "username")
	}).Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
		return re.db.Select("vehicle_id", "name", "rating")
	}).Order("created_at desc").Where("vehicle_id = ?", id).Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak ditemukan")
	}

	return data, nil
}
