package vehicles

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type vehicle_service struct {
	repo interfaces.VehicleRepo
}

func NewService(repo interfaces.VehicleRepo) *vehicle_service {
	return &vehicle_service{repo}
}

func (s *vehicle_service) GetAllVehicles() *helpers.Response {
	result, err := s.repo.FindAllVehicles()
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}
	return helpers.New(result, 200, false)
}

func (s *vehicle_service) AddVehicle(data *models.Vehicle) *helpers.Response {
	result, err := s.repo.SaveVehicle(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 201, false)
}

func (s *vehicle_service) UpdateVehicle(id int, data *models.Vehicle) *helpers.Response {
	result, err := s.repo.ChangeVehicle(id, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *vehicle_service) DeleteVehicle(id int, data *models.Vehicle) *helpers.Response {
	result, err := s.repo.RemoveVehicle(id, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *vehicle_service) GetVehicleName(name string, data *models.Vehicles) *helpers.Response {
	result, err := s.repo.FindVehicleName(name, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *vehicle_service) SortByLocation(location string, data *models.Vehicles) *helpers.Response {
	result, err := s.repo.SortByLocation(location, data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	return helpers.New(result, 200, false)
}

func (s *vehicle_service) PopularVehicle() *helpers.Response {
	result, err := s.repo.PopularVehicle()
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	return helpers.New(result, 200, false)
}
