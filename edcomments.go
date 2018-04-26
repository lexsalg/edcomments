package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/lexsalg/edcomments/commons"
	"github.com/lexsalg/edcomments/migration"
	"github.com/lexsalg/edcomments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.IntVar(&commons.Port, "port", 8000, "Puerto para el server")
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
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Run server http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa(server)")

}
