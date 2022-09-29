package vehicles

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/gorental/src/database/orm/models"
)

func TestGetAllVehicles(t *testing.T) {
	repo := RepoVehicleMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Vehicles{
		{VehicleId: 1, Name: "test 1"},
		{VehicleId: 2, Name: "test 2"},
	}

	repo.mock.On("FindAllVehicles").Return(&dataMock, nil)
	data := service.GetAllVehicles()

	result := data.Data.(*models.Vehicles)
	for i, v := range *result {
		assert.Equal(t, dataMock[i].VehicleId, v.VehicleId, "expect id from data mock")
	}
}

func TestAddVehicle(t *testing.T) {
	repo := RepoVehicleMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Vehicle{VehicleId: 1, Name: "test 1"}

	repo.mock.On("SaveVehicle", &dataMock).Return(&dataMock, nil)
	data := service.AddVehicle(&dataMock)

	var expectId uint = 1
	result := data.Data.(*models.Vehicle)

	assert.Equal(t, expectId, result.VehicleId, "id must be 1")
}
