package main

import (
	"flag"

	service "github.com/Alang0r/vypolnyator/pkg/service"
	storage "github.com/Alang0r/vypolnyator/pkg/storage"
	_ "github.com/Alang0r/vypolnyator/sklad/api"
)

func main() {

	s := storage.NewPGStorage()
	listenAddr := flag.String("listenaddr", ":3002", "listening address")
	flag.Parse()

	srv := service.NewService("sklad", *listenAddr, &s)
	srv.SetDB()
	srv.Listen()
}
