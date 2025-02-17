package service_prod

import "database/sql"

type ServiceProd struct {
	DB *sql.DB
}

// This method is implement only to the ServiceUser
// be considered a Service
func (service *ServiceProd) New() {}
