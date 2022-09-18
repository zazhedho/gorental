package histories

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

var response helpers.Response

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewService(repo interfaces.HistoryRepo) *history_service {
	return &history_service{repo}
}

func (s *history_service) GetAllHistories() (*helpers.Response, error) {
	data, err := s.repo.FindAllHistories()
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}
	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *history_service) AddHistory(data *models.History) (*helpers.Response, error) {
	data, err := s.repo.SaveHistory(data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}
	res := response.ResponseJSON(201, data)
	return res, nil
}

func (s *history_service) UpdateHistory(r *http.Request, data *models.History) (*helpers.Response, error) {
	data, err := s.repo.ChangeHistory(r, data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}
	res := response.ResponseJSON(201, data)
	return res, nil
}

func (s *history_service) DeleteHistory(r *http.Request, data *models.History) (*helpers.Response, error) {
	data, err := s.repo.RemoveHistory(r, data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}
	res := response.ResponseJSON(200, data)
	return res, nil
}

func (s *history_service) GetHistoryByVehicleId(r *http.Request, data *models.Histories) (*helpers.Response, error) {
	data, err := s.repo.FindHistoryByVehicleId(r, data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}
	res := response.ResponseJSON(200, data)
	return res, nil
}
