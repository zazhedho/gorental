package vehicles

import (
	"fmt"
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

func TestPopularVehicle(t *testing.T) {
	repo := RepoVehicleMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Vehicles{
		{Rating: 5, Name: "test 1"},
		{Rating: 4, Name: "test 2"},
	}

	repo.mock.On("PopularVehicle").Return(&dataMock, nil)
	data := service.PopularVehicle()

	result := data.Data.(*models.Vehicles)
	for i, v := range *result {
		assert.Equal(t, dataMock[i].Rating, v.Rating, "expect rating from data mock")
	}
}

func TestGetVehicleName(t *testing.T) {
	repo := RepoVehicleMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Vehicles{
		{Name: "vario"},
	}

	repo.mock.On("FindVehicleName", "vario", &dataMock).Return(&dataMock, nil)
	data := service.GetVehicleName("vario", &dataMock)

	result := data.Data.(*models.Vehicles)
	fmt.Println(result)
	for _, v := range *result {
		assert.Equal(t, "vario", v.Name, "expect name from dataMock")
	}
}

func TestSortByLocation(t *testing.T) {
	repo := RepoVehicleMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Vehicles{
		{Location: "malang"},
	}

	repo.mock.On("SortByLocation", "malang", &dataMock).Return(&dataMock, nil)
	data := service.SortByLocation("malang", &dataMock)

	result := data.Data.(*models.Vehicles)
	fmt.Println(result)
	for _, v := range *result {
		assert.Equal(t, "malang", v.Location, "expect location from dataMock")
	}
}
