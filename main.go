package main

import (
	"log"
	"mirae-code/cmd"
	"mirae-code/env"
	"mirae-code/infra/database"
	"os"
)

func main() {
	di, err := env.NewENV("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if di.DB, di.Err = database.NewMySQLDB(di.Params); di.Err != nil {
		// Handle with middleware here upon error
		log.Fatal(di.Err)
	}
	cli := cmd.NewCLI(di, os.Args)
	if cli.Start(); cli.Error() != nil {
		log.Fatal(cli.Error())
	}
}
