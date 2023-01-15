package main

import (
	"flag"
	"log"

	service "github.com/Alang0r/vypolnyator/pkg/service"
	storage "github.com/Alang0r/vypolnyator/pkg/storage"
	_ "github.com/Alang0r/vypolnyator/sklad/api"
)

func main() {

	mem := storage.NewMemoryStorage()
	listenAddr := flag.String("listenaddr", ":3001", "listening address")
	flag.Parse()

	srv := service.NewService("sklad", *listenAddr, mem)

	srv.Start()
	log.Printf("Sklad is listening on port: %s", *listenAddr)
	// if err := srv.Start(); err != nil {
	// 	log.Fatalf("Error during startup: %s", err.Error())

	// }
}
