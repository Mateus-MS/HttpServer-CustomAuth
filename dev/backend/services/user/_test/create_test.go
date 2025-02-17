package user_test

import (
	"testing"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/app"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
	test_utils "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/test/utils"
)

func TestCreate(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	app := app.Application{
		DB: db,
	}
	app.Services = make(map[string]services.Service)
	app.Services["user"] = &service_user.ServiceUser{DB: app.DB}

	user := models.User{
		Username: "test",
		Password: "test",
		Email:    "test",
	}

	err := app.Services["user"].Create(&user)

	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}
}

func TestCreate_EmptyFields(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	app := app.Application{
		DB: db,
	}
	app.Services = make(map[string]services.Service)
	app.Services["user"] = &service_user.ServiceUser{DB: app.DB}

	user := models.User{
		Username: "",
		Password: "",
		Email:    "",
	}

	err := app.Services["user"].Create(&user)

	if err == nil {
		t.Errorf("Expected error, got nil: %v", err)
	}
}

func TestCreate_InvalidFields(t *testing.T) {

}

func TestCreate_DuplicatedUser(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	app := app.Application{
		DB: db,
	}
	app.Services = make(map[string]services.Service)
	app.Services["user"] = &service_user.ServiceUser{DB: app.DB}

	user := models.User{
		Username: "test",
		Password: "test",
		Email:    "test",
	}

	app.Services["user"].Create(&user)

	err := app.Services["user"].Create(&user)

	if err != nil {
		t.Errorf("Expected error, got nil: %v", err)
	}
}
