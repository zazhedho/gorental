package histories

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type history_ctrl struct {
	svc interfaces.HistoryService
}

func NewCtrl(ctrl interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{svc: ctrl}
}

func (c *history_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllHistories().Send(w)
}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {

	var data models.History
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}
	c.svc.AddHistory(&data).Send(w)
}

func (c *history_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {

	var data models.History
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
	c.svc.UpdateHistory(id, &data).Send(w)

}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {

	var data models.History
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	c.svc.DeleteHistory(id, &data).Send(w)
}

func (c *history_ctrl) GetHistoryByVehicleId(w http.ResponseWriter, r *http.Request) {
	var data models.Histories

	vars := r.URL.Query().Get("vehicle_id")
	id, err := strconv.Atoi(string(vars))
	if err != nil {
		helpers.New(err, 500, true)
	}

	c.svc.GetHistoryByVehicleId(id, &data).Send(w)
}
