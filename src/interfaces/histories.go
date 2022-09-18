package interfaces

import (
	"net/http"

	"github.com/zazhedho/gorental/src/database/orm/models"
	"github.com/zazhedho/gorental/src/helpers"
)

type HistoryRepo interface {
	FindAllHistories() (*models.Histories, error)
	SaveHistory(data *models.History) (*models.History, error)
	ChangeHistory(r *http.Request, data *models.History) (*models.History, error)
	RemoveHistory(r *http.Request, data *models.History) (*models.History, error)
	FindHistoryByVehicleId(r *http.Request, data *models.Histories) (*models.Histories, error)
}

type HistoryService interface {
	GetAllHistories() (*helpers.Response, error)
	AddHistory(data *models.History) (*helpers.Response, error)
	UpdateHistory(r *http.Request, data *models.History) (*helpers.Response, error)
	DeleteHistory(r *http.Request, data *models.History) (*helpers.Response, error)
	GetHistoryByVehicleId(r *http.Request, data *models.Histories) (*helpers.Response, error)
}
