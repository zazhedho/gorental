package histories

import (
	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	// awalan jalur
	route := rt.PathPrefix("/api/v1/histories").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", ctrl.GetAllHistories).Methods("GET")
	route.HandleFunc("/search", ctrl.GetHistoryByVehicleId).Methods("GET")
	route.HandleFunc("", middleware.MultipleMiddleware(ctrl.AddHistory, "admin", middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/{id}", middleware.MultipleMiddleware(ctrl.UpdateHistory, "admin", middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.MultipleMiddleware(ctrl.DeleteHistory, "admin", middleware.CheckAuth)).Methods("DELETE")
}
