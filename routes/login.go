package routes

import (
	"github.com/gorilla/mux"
	"github.com/lexsalg/edcomments/controllers"
)

//SetLoginRouter router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
