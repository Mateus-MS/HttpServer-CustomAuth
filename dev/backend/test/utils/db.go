package test_utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func SetupTestDB(t *testing.T) *sql.DB {
	// Get the project root folder
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to get module root: %v", err)
	}

	// Trim the output to remove the newline character
	moduleRoot := strings.TrimSpace(string(output))
	// Insert the .env file path
	envPath := filepath.Join(moduleRoot, "dev/.env")

	// Load the .env file
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=testestdb sslmode=disable port=42069", os.Getenv("DBuser"), os.Getenv("DBpass")))
	if err != nil {
		t.Fatalf("Failed to connect to testdb: %v", err)
	}

	// Migrations
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tb_user (
			id SERIAL PRIMARY KEY,
			
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email    VARCHAR(255) NOT NULL,

			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	return db
}

func TeardownTestDB(t *testing.T, db *sql.DB) {
	if db != nil {
		_, err := db.Exec("DROP TABLE IF EXISTS tb_user")
		if err != nil {
			t.Fatalf("Failed to drop table: %v", err)
		}
		db.Close()
	}
}
