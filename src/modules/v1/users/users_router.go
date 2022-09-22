package users

import (
	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	// awalan jalur
	route := rt.PathPrefix("/api/v1/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", ctrl.GetAllUsers).Methods("GET")
	route.HandleFunc("/profile", middleware.CheckAuth(ctrl.GetUser)).Methods("GET")
	route.HandleFunc("", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/profile", middleware.CheckAuth(ctrl.UpdateUser)).Methods("PUT")
	route.HandleFunc("/{username}", middleware.CheckAuth(ctrl.DeleteUser)).Methods("DELETE")
}
