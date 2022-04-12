package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

//Template parse
var temp = template.Must(template.ParseGlob("templates/*.html"))

//By now loads .env and execute index
func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

//Build index page
func index(w http.ResponseWriter, r *http.Request) {
	db := databaseConect()

	//Database Query
	producsSelect, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	//Get rows from select and instance on p, following appends on products
	for producsSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = producsSelect.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}

func databaseConect() *sql.DB {
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
