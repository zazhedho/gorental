package vehicles

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

var response helpers.Response

type vehicle_service struct {
	repo interfaces.VehicleRepo
}

func NewService(repo interfaces.VehicleRepo) *vehicle_service {
	return &vehicle_service{repo}
}

func (s *vehicle_service) GetAllVehicles() (*helpers.Response, error) {
	data, err := s.repo.FindAllVehicles()
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *vehicle_service) AddVehicle(data *models.Vehicle) (*helpers.Response, error) {
	data, err := s.repo.SaveVehicle(data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *vehicle_service) UpdateVehicle(r *http.Request, data *models.Vehicle) (*helpers.Response, error) {
	data, err := s.repo.ChangeVehicle(r, data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *vehicle_service) DeleteVehicle(r *http.Request, data *models.Vehicle) (*helpers.Response, error) {
	data, err := s.repo.RemoveVehicle(r, data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *vehicle_service) GetVehicleName(r *http.Request, data *models.Vehicles) (*helpers.Response, error) {
	data, err := s.repo.FindVehicleName(r, data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *vehicle_service) PopularVehicle() (*helpers.Response, error) {
	data, err := s.repo.PopularVehicle()
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, data)
	return res, nil
}
