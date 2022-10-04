package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

var repos = RepoVehicleMock{mock.Mock{}}
var service = NewService(&repos)
var ctrl = NewCtrl(service)

func TestCtrlGetAllVehicles(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMock = models.Vehicles{
		{VehicleId: 1, Name: "test 1"},
		{VehicleId: 2, Name: "test 2"},
	}

	repos.mock.On("FindAllVehicles").Return(&dataMock, nil)
	mux.HandleFunc("/test/vehicles", ctrl.GetAllVehicles)

	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/vehicles", nil))

	var vehicle models.Vehicles

	response := helpers.Response{
		Data: &vehicle,
	}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	fmt.Println(response)

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, response.IsError, "isError must be false")
}

// func TestCtrlAddVehicle(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	mux := mux.NewRouter()

// 	dataMock := models.Vehicle{
// 		VehicleId:   0,
// 		Name:        "",
// 		Location:    "",
// 		Description: "",
// 		Price:       0,
// 		Status:      "",
// 		Stock:       0,
// 		Category:    "",
// 		Rating:      0,
// 		Image:       "",
// 		CreatedAt:   time.Time{},
// 		UpdatedAt:   time.Time{},
// 	}

// 	repos.mock.On("SaveVehicle").Return(&dataMock, nil)
// 	mux.HandleFunc("/test/vehicles", ctrl.AddVehicle)

// 	mux.ServeHTTP(w, httptest.NewRequest("POST", "/test/vehicles", r))

// 	var vehicle models.Vehicle

// 	response := helpers.Response{
// 		Data: &vehicle,
// 	}

// 	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(response)

// 	assert.Equal(t, 200, w.Code, "status code must be 200")
// 	assert.False(t, response.IsError, "isError must be false")
// }
