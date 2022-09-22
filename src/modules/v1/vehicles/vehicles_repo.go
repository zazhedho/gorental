package vehicles

import (
	"errors"

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
		return nil, errors.New("tidak dapat menampilkan data")
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

func (re *vehicle_repo) ChangeVehicle(id int, data *models.Vehicle) (*models.Vehicle, error) {

	result := re.db.Model(&data).Where("vehicle_id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data kendaraan tidak ditemukan")
	}

	return data, nil
}

func (re *vehicle_repo) RemoveVehicle(id int, data *models.Vehicle) (*models.Vehicle, error) {
	result := re.db.Delete(data, id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data kendaraan tidak ditemukan")
	}

	return nil, nil
}

// Search by vehicle name
func (re *vehicle_repo) FindVehicleName(name string, data *models.Vehicles) (*models.Vehicles, error) {

	result := re.db.Order("created_at desc").Where("LOWER(name) LIKE ?", "%"+name+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	if result.RowsAffected == 1 {
		return nil, errors.New("data kendaraan tidak ditemukan")
	}

	return data, nil
}

// Search by location
func (re *vehicle_repo) SortByLocation(location string, data *models.Vehicles) (*models.Vehicles, error) {

	result := re.db.Order("rating desc").Where("LOWER(location) LIKE ?", "%"+location+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("lokasi tidak ditemukan")
	}

	return data, nil
}

func (re *vehicle_repo) PopularVehicle() (*models.Vehicles, error) {

	var data models.Vehicles
	result := re.db.Order("rating desc").Find(&data)
	if result.Error != nil {
		return nil, errors.New("tidak dapat menampilkan data")
	}

	return &data, nil
}
