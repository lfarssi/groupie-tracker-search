package main

import (
	routes "groupie_tracker/Routes"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		os.Exit(1)
	}
		routes.Router()
}