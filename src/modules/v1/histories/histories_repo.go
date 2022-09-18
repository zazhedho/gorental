package histories

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
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

	result := re.db.Preload("User").Preload("Vehicle").Order("created_at desc").Find(&data)

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

	return nil, nil
}

func (re *history_repo) ChangeHistory(r *http.Request, data *models.History) (*models.History, error) {
	vars := mux.Vars(r)
	result := re.db.Model(&data).Where("history_id = ?", vars["history_id"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("id history tidak ditemukan")
	}

	return nil, nil
}

func (re *history_repo) RemoveHistory(r *http.Request, data *models.History) (*models.History, error) {
	vars := mux.Vars(r)
	result := re.db.Delete(data, vars["history_id"])

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("id history tidak ditemukan")
	}

	return nil, nil
}

// Search
func (re *history_repo) FindHistoryByVehicleId(r *http.Request, data *models.Histories) (*models.Histories, error) {

	vars := r.URL.Query().Get("vehicle_id")
	result := re.db.Preload("User").Preload("Vehicle").Order("created_at desc").Where("vehicle_id = ?", vars).Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return data, nil
}
