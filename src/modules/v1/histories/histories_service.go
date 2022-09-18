package histories

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/interfaces"
)

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewService(repo interfaces.HistoryRepo) *history_service {
	return &history_service{repo}
}

func (s *history_service) GetAllHistories() (*models.Histories, error) {
	data, err := s.repo.FindAllHistories()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) AddHistory(data *models.History) (*models.History, error) {
	data, err := s.repo.SaveHistory(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) UpdateHistory(r *http.Request, data *models.History) (*models.History, error) {
	data, err := s.repo.ChangeHistory(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) DeleteHistory(r *http.Request, data *models.History) (*models.History, error) {
	data, err := s.repo.RemoveHistory(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *history_service) GetHistoryByVehicleId(r *http.Request, data *models.Histories) (*models.Histories, error) {
	data, err := s.repo.FindHistoryByVehicleId(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
