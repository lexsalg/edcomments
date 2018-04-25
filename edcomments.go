package main

import (
	"flag"
	"log"

	"github.com/lexsalg/edcomments/migration"
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
}
