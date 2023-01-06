package main

import (
	"flag"
	"fmt"
	"log"

	srvlib "github.com/Alang0r/vypolnyator/pkg/service"
	storage "github.com/Alang0r/vypolnyator/pkg/storage"
	telegram "github.com/Alang0r/vypolnyator/pkg/telegram"
)

func main() {

	mem := storage.NewMemoryStorage()
	listenAddr := flag.String("listenaddr", ":3001", "listening address")
	flag.Parse()

	srv := srvlib.NewService("Pechkin", *listenAddr, mem)
	srv.GetParameters(telegram.ParamTgToken)

	go srv.Start()
	log.Printf("Pechkin is listening on port: %s", *listenAddr)

	c, err := telegram.NewCommunicator(srv.Params)
	if err != nil {
		fmt.Errorf("Error start communicator: %s", err)
	}

	c.Listen()
	c.Start()

}
