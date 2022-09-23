package vehicles

import (
	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/popular", ctrl.PopularVehicle).Methods("GET")
	route.HandleFunc("/search", ctrl.GetVehicleName).Methods("GET")
	route.HandleFunc("/sort", ctrl.SortByLocation).Methods("GET")
	route.HandleFunc("", ctrl.GetAllVehicles).Methods("GET")
	route.HandleFunc("", middleware.MultipleMiddleware(ctrl.AddVehicle, "admin", middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/{id}", middleware.MultipleMiddleware(ctrl.UpdateVehicle, "admin", middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.MultipleMiddleware(ctrl.DeleteVehicle, "admin", middleware.CheckAuth)).Methods("DELETE")
}
