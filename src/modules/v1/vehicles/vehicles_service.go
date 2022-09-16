package vehicles

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/interfaces"
)

type vehicle_service struct {
	repo interfaces.VehicleRepo
}

func NewService(repo interfaces.VehicleRepo) *vehicle_service {
	return &vehicle_service{repo}
}

func (s *vehicle_service) GetAllVehicles() (*models.Vehicles, error) {
	data, err := s.repo.FindAllVehicles()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *vehicle_service) AddVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	data, err := s.repo.SaveVehicle(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *vehicle_service) UpdateVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	data, err := s.repo.ChangeVehicle(r, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *vehicle_service) DeleteVehicle(r *http.Request, data *models.Vehicle) (*models.Vehicle, error) {
	data, err := s.repo.RemoveVehicle(r, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
