package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/Alang0r/vypolnyator/pechkin/internal/handlers"
	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/Alang0r/vypolnyator/pkg/telegram"
)

func main() {

	mem := storage.NewMemoryStorage()
	listenAddr := flag.String("listenaddr", ":3002", "listening address")
	flag.Parse()

	srv := service.NewService("Pechkin", *listenAddr, mem)
	srv.GetParameters(telegram.ParamTgToken)

	go srv.Start()
	log.Printf("Pechkin is listening on port: %s", *listenAddr)

	c, err := telegram.NewCommunicator(srv.Params)
	if err.Code != error.ErrCodeNone {
		fmt.Errorf("Error start communicator: %s", err)
	}

	c.Listen()
	c.Start()

}
