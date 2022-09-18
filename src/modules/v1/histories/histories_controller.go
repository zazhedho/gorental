package histories

import (
	"encoding/json"
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/interfaces"
)

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
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := c.svc.AddHistory(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (c *history_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data, err := c.svc.UpdateHistory(r, &datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.History

	data, err := c.svc.DeleteHistory(r, &datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (c *history_ctrl) GetHistoryByVehicleId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.Histories

	data, err := c.svc.GetHistoryByVehicleId(r, &datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}
