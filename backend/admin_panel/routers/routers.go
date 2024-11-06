package routers

import (
	"admin_panel/controllers"

	"github.com/gorilla/mux"
)

func RegisterRouters() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/admin/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("admin/products", controllers.CreateProduct).Methods("POST")

	return router
}
