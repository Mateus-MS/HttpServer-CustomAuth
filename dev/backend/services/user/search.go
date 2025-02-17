package service_user

import (
	"database/sql"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
)

func (service *ServiceUser) Search(userOBJ *models.User) (user *models.User, err error) {
	query := `
		SELECT
			id,
			username,
			email
		FROM tb_user
		WHERE username = $1
		OR email = $2
	`

	err = service.DB.QueryRow(query, userOBJ.Username, userOBJ.Email).Scan(&userOBJ.ID, &userOBJ.Username, &userOBJ.Email)
	if err == sql.ErrNoRows {
		return nil, models.ErrorInvalidUser
	} else if err != nil {
		return nil, err
	}

	return userOBJ, nil
}
