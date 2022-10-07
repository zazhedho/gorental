package auth

import (
	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/modules/v1/users"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1").Subrouter()

	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/login", ctrl.SignIn).Methods("POST")
}
