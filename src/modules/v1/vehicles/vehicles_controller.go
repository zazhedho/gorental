package vehicles

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type vehicle_ctrl struct {
	svc interfaces.VehicleService
}

func NewCtrl(ctrl interfaces.VehicleService) *vehicle_ctrl {
	return &vehicle_ctrl{ctrl}
}

func (c *vehicle_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllVehicles().Send(w)
}

func (c *vehicle_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {

	var data models.Vehicle
	data.Image = r.Context().Value("imageName").(string)

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}
	c.svc.AddVehicle(&data).Send(w)
}

func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {

	var data models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	c.svc.UpdateVehicle(id, &data).Send(w)
}

func (c *vehicle_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	var data models.Vehicle

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	c.svc.DeleteVehicle(id, &data).Send(w)
}

func (c *vehicle_ctrl) GetVehicleName(w http.ResponseWriter, r *http.Request) {

	var data models.Vehicles
	name := strings.ToLower(r.URL.Query().Get("name"))
	c.svc.GetVehicleName(name, &data).Send(w)
}

func (c *vehicle_ctrl) SortByLocation(w http.ResponseWriter, r *http.Request) {

	var data models.Vehicles
	location := strings.ToLower(r.URL.Query().Get("location"))
	c.svc.SortByLocation(location, &data).Send(w)
}

func (c *vehicle_ctrl) PopularVehicle(w http.ResponseWriter, r *http.Request) {
	c.svc.PopularVehicle().Send(w)
}

func (c *vehicle_ctrl) SortByType(w http.ResponseWriter, r *http.Request) {

	var data models.Vehicles
	category := strings.ToLower(r.URL.Query().Get("category"))
	c.svc.SortByType(category, &data).Send(w)
}
