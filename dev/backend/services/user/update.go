package service_user

import (
	"database/sql"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
)

func Update(baseUser *models.User, newUser *models.User, db *sql.DB) error {
	if err := newUser.Validate(); err != nil {
		return err
	}

	if _, err := Search(baseUser, db); err != nil {
		println(err.Error())
		return models.ErrorUserNotFound
	}

	sessionToken := newUser.SessionToken

	//Try update the values from the user
	if _, err := db.Exec(
		`UPDATE tb_user SET session_token = $1 WHERE username = $2`,
		sessionToken, baseUser.Username); err != nil {
		return err
	}

	return nil
}
