package main

import (
	"admin_panel/config"
	"admin_panel/routers"
	"log"
	"net/http"
)

func main() {
	config.InitDB()

	router := routers.RegisterRouters()

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
