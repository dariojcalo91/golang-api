package main

import (
	"api/internal/adapters"
	"log"
)

func main() {

	server := adapters.NewServer()
	if err := server.Run(":3000"); err != nil {
		log.Fatalln("Can't start the server")
	}
}
