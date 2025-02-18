package service_user_test

import (
	"testing"

	_ "github.com/lib/pq"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
	test_utils "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/test/utils"
)

func TestCreate(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	user := models.User{
		Username: "test",
		Password: "test",
		Email:    "test",
	}

	err := service_user.Create(&user, db)

	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}
}

func TestCreate_EmptyFields(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	user := models.User{
		Username: "",
		Password: "",
		Email:    "",
	}

	err := service_user.Create(&user, db)

	if err == nil {
		t.Errorf("Expected error, got nil: %v", err)
	}
}

func TestCreate_DuplicatedUser(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	user := models.User{
		Username: "test",
		Password: "test",
		Email:    "test",
	}

	service_user.Create(&user, db)

	err := service_user.Create(&user, db)

	if err == nil {
		t.Errorf("Expected error, got nil: %v", err)
	}
}
