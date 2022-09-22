package interfaces

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

type VehicleRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	SaveVehicle(data *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(id int, data *models.Vehicle) (*models.Vehicle, error)
	RemoveVehicle(id int, data *models.Vehicle) (*models.Vehicle, error)
	FindVehicleName(name string, data *models.Vehicles) (*models.Vehicles, error)
	SortByLocation(location string, data *models.Vehicles) (*models.Vehicles, error)
	PopularVehicle() (*models.Vehicles, error)
}

type VehicleService interface {
	GetAllVehicles() *helpers.Response
	AddVehicle(data *models.Vehicle) *helpers.Response
	UpdateVehicle(id int, data *models.Vehicle) *helpers.Response
	DeleteVehicle(id int, data *models.Vehicle) *helpers.Response
	GetVehicleName(name string, data *models.Vehicles) *helpers.Response
	SortByLocation(location string, data *models.Vehicles) *helpers.Response
	PopularVehicle() *helpers.Response
}
