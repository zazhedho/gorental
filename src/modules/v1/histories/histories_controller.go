package histories

import (
	"encoding/json"
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

var response helpers.Response

type history_ctrl struct {
	svc interfaces.HistoryService
}

func NewCtrl(repo interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{svc: repo}
}

func (c *history_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := c.svc.GetAllHistories()
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat mengambil data", data, err)
	} else {
		response.ResponseJSON(w, 200, "success", data, err)
	}
}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		data, err := c.svc.AddHistory(&datas)
		if err != nil {
			response.ResponseJSON(w, 400, "Tidak dapat menambahkan data", data, err)
		} else {
			response.ResponseJSON(w, 200, "success", data, err)
		}
	}
}

func (c *history_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		data, err := c.svc.UpdateHistory(r, &datas)
		if err != nil {
			response.ResponseJSON(w, 400, "Tidak dapat mengubah data", data, err)
		} else {
			response.ResponseJSON(w, 200, "success", data, err)
		}
	}

}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	data, err := c.svc.DeleteHistory(r, &datas)
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menghapus data", data, err)
	} else {
		response.ResponseJSON(w, 200, "success", data, err)
	}
}

func (c *history_ctrl) GetHistoryByVehicleId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.Histories
	data, err := c.svc.GetHistoryByVehicleId(r, &datas)
	if err != nil {
		response.ResponseJSON(w, 400, "Tidak dapat menampilkan data", data, err)
	} else {
		response.ResponseJSON(w, 200, "success", data, err)
	}
}
