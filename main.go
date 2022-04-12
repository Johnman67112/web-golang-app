package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "T-Shirt", Description: "Blue, really pretty", Price: 39, Quantity: 5},
		{"Sneakers", "Comfortable", 89, 10},
		{"Headphone", "Very good", 59, 2},
		{"New Product", "Really cool", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
