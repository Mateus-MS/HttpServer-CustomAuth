package service_user

import (
	"database/sql"
)

type ServiceUser struct {
	DB *sql.DB
}

// This method is implement only to the ServiceUser
// be considered a Service
func (service *ServiceUser) New() {}
