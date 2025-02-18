package service_user_test

import (
	"testing"

	_ "github.com/lib/pq"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
	test_utils "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/test/utils"
)

func TestSearch(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	user := models.User{
		Username: "test",
		Email:    "test",
		Password: "test",
	}

	service_user.Create(&user, db)

	userOBJ, err := service_user.Search(&user, db)

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

	user := models.User{
		Username: "test",
		Email:    "test",
		Password: "test",
	}

	_, err := service_user.Search(&user, db)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
