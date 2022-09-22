package histories

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
	"github.com/zazhedho/gorental/src/interfaces"
)

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewService(repo interfaces.HistoryRepo) *history_service {
	return &history_service{repo}
}

func (s *history_service) GetAllHistories() *helpers.Response {
	result, err := s.repo.FindAllHistories()
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *history_service) AddHistory(data *models.History) *helpers.Response {
	result, err := s.repo.SaveHistory(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 201, false)
}

func (s *history_service) UpdateHistory(id int, data *models.History) *helpers.Response {
	result, err := s.repo.ChangeHistory(id, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *history_service) DeleteHistory(id int, data *models.History) *helpers.Response {
	result, err := s.repo.RemoveHistory(id, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *history_service) GetHistoryByVehicleId(id int, data *models.Histories) *helpers.Response {
	result, err := s.repo.FindHistoryByVehicleId(id, data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	return helpers.New(result, 200, false)
}
