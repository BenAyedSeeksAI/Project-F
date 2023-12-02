package main

import (
	"fmt"
	"os"
	"rhmanager/db"
	"rhmanager/routing"
)

func TerminalInput() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "DbSeed" {
			db.EstablishDB()
			fmt.Println("¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤")
			fmt.Println("DB has been established...")
			fmt.Println("¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤")
		} else {
			fmt.Println("Hello Project")
		}
	} else {
		fmt.Println("Hello Project Noob")
	}

}
func PrepareEndpoints() {
	routing.RunRouting()
}
func main() {
	TerminalInput()
	PrepareEndpoints()
}
