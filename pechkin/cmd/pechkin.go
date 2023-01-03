package main

import (
	"flag"
	"log"

	service "github.com/Alang0r/vypolnyator/pkg/service"
	storage "github.com/Alang0r/vypolnyator/pkg/storage"
)

func main() {

	mem := storage.NewMemoryStorage()
	listenAddr := flag.String("listenaddr", ":3001", "listening address")
	flag.Parse()

	srv := service.NewService("Pechkin", *listenAddr, mem)

	srv.Start()
	log.Printf("Pechkin is listening on port: %s", *listenAddr)
	// if err := srv.Start(); err != nil {
	// 	log.Fatalf("Error during startup: %s", err.Error())

	// }
}
