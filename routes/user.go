package routes

import (
	"github.com/gorilla/mux"
	"github.com/lexsalg/edcomments/controllers"
	"github.com/urfave/negroni"
)

// SetUserRouter ruta para el registro de usuario
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
