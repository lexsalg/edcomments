package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/lexsalg/edcomments/migration"
	"github.com/lexsalg/edcomments/routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración de la BD...")
		migration.Migrate()
		log.Println("Migración de la BD terminada.")
	}

	// inicia las rutas
	router := routes.InitRoutes()

	//inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    ":8000",
		Handler: n,
	}

	log.Println("Run server http://localhost:8000")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa(server)")

}
