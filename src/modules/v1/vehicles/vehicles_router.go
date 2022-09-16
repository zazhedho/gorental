package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAllVehicles).Methods("GET")
	route.HandleFunc("/", ctrl.AddVehicle).Methods("POST")
	route.HandleFunc("/{vehicle_id}", ctrl.UpdateVehicle).Methods("PUT")
	route.HandleFunc("/{vehicle_id}", ctrl.DeleteVehicle).Methods("DELETE")
}
