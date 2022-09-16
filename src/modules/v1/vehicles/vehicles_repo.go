package vehicles

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/database/orm/models"
	"gorm.io/gorm"
)

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicle_repo {
	return &vehicle_repo{db}
}

func (re *vehicle_repo) FindAllVehicles() (*models.Vehicles, error) {
	var data models.Vehicles

	result := re.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (re *vehicle_repo) SaveVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	result := re.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambah data")
	}

	return data, nil
}

func (re *vehicle_repo) ChangeVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	vars := mux.Vars(r)
	result := re.db.Model(&data).Where("vehicle_id = ?", vars["vehicle_id"]).Updates(data)

	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	return data, nil
}

func (re *vehicle_repo) RemoveVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	vars := mux.Vars(r)
	result := re.db.Delete(data, vars["vehicle_id"])

	if result.Error != nil {
		return nil, errors.New("gagal hapus data")
	}

	return data, nil
}
