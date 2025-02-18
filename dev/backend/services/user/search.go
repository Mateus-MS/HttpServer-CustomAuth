package service_user

import (
	"database/sql"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
)

func Search(userOBJ *models.User, db *sql.DB) (user *models.User, err error) {
	query := `
		SELECT
			id,
			username,
			email,
			session_token
		FROM tb_user
		WHERE username = $1
		OR email = $2
	`

	println("Searching for user: ", userOBJ.Username)

	err = db.QueryRow(query, userOBJ.Username, userOBJ.Email).Scan(&userOBJ.ID, &userOBJ.Username, &userOBJ.Email, &userOBJ.SessionToken)
	if err == sql.ErrNoRows {
		return nil, models.ErrorInvalidUser
	} else if err != nil {
		return nil, err
	}

	return userOBJ, nil
}
