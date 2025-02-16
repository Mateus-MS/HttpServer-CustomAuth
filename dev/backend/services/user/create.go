package service_user

import "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"

func (service *ServiceUser) Create(userOBJ *models.User) error {
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

	if !validateUsername(userOBJ.Username) {
		return models.ErrorInvalidUser
	}

	if !validatePassword(userOBJ.Password) {
		return models.ErrorInvalidUser
	}

	if !validateEmail(userOBJ.Email) {
		return models.ErrorInvalidUser
	}

	if _, err := service.DB.Exec(query, userOBJ.Username, userOBJ.Password, userOBJ.Email); err != nil {
		return err
	}

	return nil
}

func validateUsername(username string) bool {
	return username != ""
}

func validatePassword(password string) bool {
	return password != ""
}

func validateEmail(email string) bool {
	return email != ""
}
