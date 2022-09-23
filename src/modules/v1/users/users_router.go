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

	route.HandleFunc("", middleware.MultipleMiddleware(ctrl.GetAllUsers, "admin", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/profile", middleware.MultipleMiddleware(ctrl.GetUser, "user", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/profile", middleware.MultipleMiddleware(ctrl.UpdateUser, "user", middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{username}", middleware.MultipleMiddleware(ctrl.DeleteUser, "admin", middleware.CheckAuth)).Methods("DELETE")
}
