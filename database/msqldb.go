package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectMySQL(host string, user string, password string, dbName string, port int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
