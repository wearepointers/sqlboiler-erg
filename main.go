package main

import (
	"github.com/expanse-agency/sqlboiler-erg/utils"
)

func main() {
	c, err := utils.ParseConfig()
	if err != nil {
		panic(err)
	}

	if err := c.ConvertSQLBoilerModelsToApiModels(); err != nil {
		panic(err)
	}
}
