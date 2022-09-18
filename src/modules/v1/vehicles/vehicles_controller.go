package vehicles

import (
	"encoding/json"
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

var response helpers.Response

type vehicle_ctrl struct {
	svc interfaces.VehicleService
}

func NewCtrl(repo interfaces.VehicleService) *vehicle_ctrl {
	return &vehicle_ctrl{svc: repo}
}

func (c *vehicle_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := c.svc.GetAllVehicles()
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menampilkan data", data, err)
	} else {
		response.ResponseJSON(w, 200, "Success", data, err)
	}
}

func (c *vehicle_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		data, err := c.svc.AddVehicle(&datas)
		if err != nil {
			response.ResponseJSON(w, 400, "Tidak dapat menambahkan data", data, err)
		} else {
			response.ResponseJSON(w, 200, "success", data, err)
		}
	}
}

func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		data, err := c.svc.UpdateVehicle(r, &datas)
		if err != nil {
			response.ResponseJSON(w, 400, "Tidak dapat mengubah data", data, err)
		} else {
			response.ResponseJSON(w, 200, "success", data, err)
		}
	}

}

func (c *vehicle_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.Vehicle

	data, err := c.svc.DeleteVehicle(r, &datas)
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menghapus data", data, err)
	} else {
		response.ResponseJSON(w, 200, "success", data, err)
	}
}

func (c *vehicle_ctrl) GetVehicleName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var datas models.Vehicles

	data, err := c.svc.GetVehicleName(r, &datas)
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menampilkan data", data, err)
	} else {
		response.ResponseJSON(w, 200, "success", data, err)
	}
}

func (c *vehicle_ctrl) PopularVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := c.svc.PopularVehicle()
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menampilkan data", data, err)
	} else {
		response.ResponseJSON(w, 200, "success", data, err)
	}
}
