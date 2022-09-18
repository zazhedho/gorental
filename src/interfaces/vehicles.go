package interfaces

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

type VehicleRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	SaveVehicle(data *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
	RemoveVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error)
	FindVehicleName(r *http.Request, data *models.Vehicles) (*models.Vehicles, error)
	PopularVehicle() (*models.Vehicles, error)
}

type VehicleService interface {
	GetAllVehicles() (*helpers.Response, error)
	AddVehicle(data *models.Vehicle) (*helpers.Response, error)
	UpdateVehicle(r *http.Request, data *models.Vehicle) (*helpers.Response, error)
	DeleteVehicle(r *http.Request, data *models.Vehicle) (*helpers.Response, error)
	GetVehicleName(r *http.Request, data *models.Vehicles) (*helpers.Response, error)
	PopularVehicle() (*helpers.Response, error)
}
