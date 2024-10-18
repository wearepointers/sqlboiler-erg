package main

import (
	"log"

	"github.com/wearepointers/sqlboiler-erg/example/database"
)

func main() {
	if err := database.Migrater(); err != nil {
		log.Fatal(err)
	}
}
