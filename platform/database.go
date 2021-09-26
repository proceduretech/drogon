package platform

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	database *sql.DB
)

var once sync.Once

const (
	host   = "localhost"
	port   = 5432
	user   = "ulhas"
	dbname = "drogon"
)

func InitializeDatabase() {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
			host, port, user, dbname)

		var error error
		database, error = sql.Open("postgres", psqlInfo)

		if error != nil {
			log.Fatal(error)
		}

		error = database.Ping()

		if error != nil {
			log.Fatal(error)
		}

		log.Println("Database connection successfully made")
	})
}

func GetDatabase() *sql.DB {
	return database
}
