package main

import (
	"os"
	"rhmanager/db"
	"rhmanager/routing"
)

func CommandProcessing() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "DbSeed" {
			db.EstablishDB()
		}
	} else {
		routing.RunRouting()
	}
}

func main() {
	CommandProcessing()
}
