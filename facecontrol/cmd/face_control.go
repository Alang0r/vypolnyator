package main

import (
	"flag"

	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/storage"
)

func main() {
	s := storage.NewPGStorage()
	listenAddr := flag.String("listenaddr", ":3003", "listening address")
	flag.Parse()

	srv := service.NewService("face-control", *listenAddr, &s)
	srv.SetDB()
	srv.Listen()
}
