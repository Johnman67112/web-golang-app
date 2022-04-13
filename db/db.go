package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func DatabaseConect() *sql.DB {
	//Get envs setted on main
	user := os.Getenv("USER")
	dbase := os.Getenv("DBNAME")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	ssl := os.Getenv("SSLMODE")

	//Build conn string
	conn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbase, pass, host, ssl)

	//Connect database
	db, err := sql.Open(user, conn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
