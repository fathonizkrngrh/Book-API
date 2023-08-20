package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
    username string = "root"
    password string = "fathoni"
    database string = "quiz3"
)

var (
    dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
	counts int64
)

func MySQL() (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)

    if err != nil {
        return nil, err
    }

    return db, nil
}

func ConnectToDB() *sql.DB{
	// dsn := os.Getenv("DSN")

	for {
		connection, err := MySQL()
		if err != nil {
			log.Println("MySQL not ready....")
			counts++
		} else {
			log.Println("Connected to MySQL")
			return connection
		}

		if counts >10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}