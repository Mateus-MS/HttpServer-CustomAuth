package service_user

import (
	"database/sql"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
)

func Create(userOBJ *models.User, db *sql.DB) error {
	query := `
        INSERT INTO tb_user (
            username, 
            password,
            email
        ) VALUES (
            $1,
            $2,
            $3
        )
    `

	if err := userOBJ.Validate(); err != nil {
		return err
	}

	if _, err := db.Exec(query, userOBJ.Username, userOBJ.PasswordHash, userOBJ.Email); err != nil {
		return err
	}

	return nil
}
