package main

import (
	"database/sql"
	"fmt"
	"kafkaP/server/handlers"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "__ADD__CONN_STRING")
	if err != nil {
		panic("DB connection failed: " + err.Error())
	}

	defer db.Close()

	//Ping DB to verify connection
	if err := db.Ping(); err != nil {
		panic("DB not reachable: " + err.Error())
	} else {
		fmt.Println("DB connected successfully")
	}

	http.HandleFunc("/incidents", handlers.IncidentHandler(db))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Server failed to start: " + err.Error())
	} else {
		fmt.Println("Server started successfully on port: ", 8080)
	}
}
