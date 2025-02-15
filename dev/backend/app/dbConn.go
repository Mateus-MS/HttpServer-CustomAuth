package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

func startDBConnection(user, pass, dbname string) *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=42069", user, pass, dbname)

	client, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to postgreSQL: ", err)
	}

	return client
}

var instance *sql.DB
var once sync.Once

func GetInstance() *sql.DB {
	once.Do(func() {
		instance = startDBConnection(os.Getenv("DBuser"), os.Getenv("DBpass"), "cluster")
	})
	return instance
}
