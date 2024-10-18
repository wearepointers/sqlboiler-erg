package main

import (
	"log"

	"github.com/wearepointers/sqlboiler-erg/utils"
)

func main() {
	c, err := utils.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := c.ConvertSQLBoilerModelsToApiModels(); err != nil {
		log.Fatal(err)
	}
}
