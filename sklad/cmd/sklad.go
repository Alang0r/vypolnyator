package main

import (
	"flag"

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
}
