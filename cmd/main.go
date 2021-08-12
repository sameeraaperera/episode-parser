package main

import (
	"log"
	"os"

	"github.com/sap4001/episode-parser/internal/server"
)

func main() {
	port, err := server.GetListenPort(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}
	server := server.NewServer(port)
	log.Println("Starting HTTP Server on port: ", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: unexpected error in ListenAndServe: %v", err)
	}
}
