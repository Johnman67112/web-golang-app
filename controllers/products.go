package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Johnman67112/web_go/models"
)

//Template parse
var temp = template.Must(template.ParseGlob("templates/*.html"))

//Builds index page
func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

//Builds new page
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

//Inserts new product
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")

		sprice := r.FormValue("price")
		squantity := r.FormValue("quantity")
		price, quantity := productNumParse(sprice, squantity)

		models.SetProduct(name, description, price, quantity)
	}
	http.Redirect(w, r, "/", 301)
}

//Parse number fields from insert request
func productNumParse(price, quantity string) (float64, int) {
	convertedPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Error on price conversion")
	}

	convertedQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		log.Println("Error on quantity conversion")
	}
	return convertedPrice, convertedQuantity
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}
