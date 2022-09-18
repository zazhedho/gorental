package histories

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/histories").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAllHistories).Methods("GET")
	route.HandleFunc("/search", ctrl.GetHistoryByVehicleId).Methods("GET")
	route.HandleFunc("/", ctrl.AddHistory).Methods("POST")
	route.HandleFunc("/{history_id}", ctrl.UpdateHistory).Methods("PUT")
	route.HandleFunc("/{history_id}", ctrl.DeleteHistory).Methods("DELETE")
}
