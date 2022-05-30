package main

import (
	"log"

	"github.com/zak-blake/microprog"
)

func main() {
	err := microprog.StartProgrammableServer()
	if err != nil {
		log.Fatal(err)
	}
}
