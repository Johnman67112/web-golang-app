package main

import (
	"net/http"

	"github.com/Johnman67112/web_go/routes"
	"github.com/joho/godotenv"
)

//By now loads .env and execute index
func main() {
	godotenv.Load(".env")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
