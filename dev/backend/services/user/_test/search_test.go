package user_test

import (
	"testing"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/app"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
	test_utils "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/test/utils"
)

func TestSearch(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	app := app.Application{
		DB: db,
	}
	app.Services = make(map[string]services.Service)
	app.Services["user"] = &service_user.ServiceUser{DB: app.DB}

	user := models.User{
		Username: "test",
		Email:    "test",
		Password: "test",
	}

	app.Services["user"].Create(&user)

	userOBJ, err := app.Services["user"].Search(&user)

	if err != nil {
		t.Error("Expected nil, got error")
	}

	if userOBJ.Username != user.Username {
		t.Errorf("Expected %s, got %s", user.Username, userOBJ.Username)
	}
}

func TestSearch_Unexisting(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	app := app.Application{
		DB: db,
	}
	app.Services = make(map[string]services.Service)
	app.Services["user"] = &service_user.ServiceUser{DB: app.DB}

	user := models.User{
		Username: "test",
		Email:    "test",
		Password: "test",
	}

	_, err := app.Services["user"].Search(&user)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
