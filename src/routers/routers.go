package routers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/zazhedho/gorental/src/database/orm"
	"github.com/zazhedho/gorental/src/modules/v1/histories"
	"github.com/zazhedho/gorental/src/modules/v1/users"
	"github.com/zazhedho/gorental/src/modules/v1/vehicles"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	histories.New(mainRoute, db)

	return mainRoute, nil
}
