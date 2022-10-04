package vehicles

import (
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/gorental/src/database/orm/models"
)

type RepoVehicleMock struct {
	mock mock.Mock
}

func (m *RepoVehicleMock) FindAllVehicles() (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoVehicleMock) SaveVehicle(data *models.Vehicle) (*models.Vehicle, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoVehicleMock) ChangeVehicle(id int, data *models.Vehicle) (*models.Vehicle, error) {
	args := m.mock.Called(id, data)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoVehicleMock) RemoveVehicle(id int, data *models.Vehicle) (*models.Vehicle, error) {
	args := m.mock.Called(id, data)
	return args.Get(0).(*models.Vehicle), nil
}

func (m *RepoVehicleMock) FindVehicleName(name string, data *models.Vehicles) (*models.Vehicles, error) {
	args := m.mock.Called(name, data)
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoVehicleMock) SortByLocation(location string, data *models.Vehicles) (*models.Vehicles, error) {
	args := m.mock.Called(location, data)
	return args.Get(0).(*models.Vehicles), nil
}

func (m *RepoVehicleMock) PopularVehicle() (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}
