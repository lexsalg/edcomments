package routes

import (
	"github.com/gorilla/mux"
	"github.com/lexsalg/edcomments/controllers"
	"github.com/urfave/negroni"
)

// SetVoteRouter es la ruta para el registro de un voto
func SetVoteRouter(router *mux.Router) {
	prefix := "/api/votes"
	subrouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subrouter.HandleFunc("/", controllers.VoteRegister).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subrouter),
		),
	)
}
