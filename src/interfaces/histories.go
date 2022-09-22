package interfaces

import (
	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

type HistoryRepo interface {
	FindAllHistories() (*models.Histories, error)
	SaveHistory(data *models.History) (*models.History, error)
	ChangeHistory(id int, data *models.History) (*models.History, error)
	RemoveHistory(id int, data *models.History) (*models.History, error)
	FindHistoryByVehicleId(id int, data *models.Histories) (*models.Histories, error)
}

type HistoryService interface {
	GetAllHistories() *helpers.Response
	AddHistory(data *models.History) *helpers.Response
	UpdateHistory(id int, data *models.History) *helpers.Response
	DeleteHistory(id int, data *models.History) *helpers.Response
	GetHistoryByVehicleId(id int, data *models.Histories) *helpers.Response
}
