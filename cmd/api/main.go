package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"quiz3/config"
	"quiz3/data"
)

var webPort = "8080"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	conn := config.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to database!")
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: Router(), 
	}

	fmt.Println(fmt.Sprintf("server running at http://localhost:%s", webPort))
	
	err:= srv.ListenAndServe()
	if err != nil {log.Panic(err)}
}

