package platform

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var database *sql.DB
var once sync.Once

const (
	host   = "localhost"
	port   = 5432
	user   = "ulhas"
	dbname = "drogon"
)

func InitializeDB() {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
			host, port, user, dbname)

		db, err := sql.Open("postgres", psqlInfo)

		if err != nil {
			log.Fatal(err)
		}

		defer database.Close()
		err = db.Ping()

		if err != nil {
			log.Fatal(err)
		}

		database = db
		log.Println("Database connection successfully made")
	})
}

func GetDatabase() *sql.DB {
	return database
}
