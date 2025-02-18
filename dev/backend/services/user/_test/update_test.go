package service_user_test

import (
	"database/sql"
	"testing"

	"github.com/Mateus-MS/HttpServerGolang.git/dev/backend/models"
	service_user "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/services/user"
	test_utils "github.com/Mateus-MS/HttpServerGolang.git/dev/backend/test/utils"
)

func TestUpdate(t *testing.T) {
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

	updatedUser := user.Copy()
	updatedUser.Password = "test2"

	err = service_user.Update(&user, &updatedUser, db)

	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}

	userSearched, _ := service_user.Search(&user, db)
	if userSearched.SessionToken != updatedUser.SessionToken {
		t.Errorf("Expected %s, got %s", updatedUser.SessionToken.String, userSearched.SessionToken.String)
	}
}

func TestUpdate_Unexisting(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	user := models.User{
		Username:     "test",
		Password:     "test",
		SessionToken: sql.NullString{String: "test", Valid: true},
	}

	updatedUser := user.Copy()
	updatedUser.SessionToken = sql.NullString{String: "test2", Valid: true}

	err := service_user.Update(&user, &updatedUser, db)

	if err == nil {
		t.Errorf("Expected error, got nil: %v", err)
	}
}

func TestUpdate_Invalid(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	defer test_utils.TeardownTestDB(t, db)

	user := models.User{
		Username:     "test",
		Password:     "test",
		SessionToken: sql.NullString{String: "test", Valid: true},
	}

	service_user.Create(&user, db)

	updatedUser := user.Copy()
	updatedUser.Password = ""

	err := service_user.Update(&user, &updatedUser, db)

	if err == nil {
		t.Errorf("Expected error, got nil: %v", err)
	}
}
