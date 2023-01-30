package main

import (
	"flag"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/Alang0r/vypolnyator/pkg/telegram"
)

var teleHandlers map[string]telegram.THandler

func main() {
	mem := storage.NewMemoryStorage()
	listenAddr := flag.String("listenaddr", ":3002", "listening address")
	flag.Parse()

	Srv := service.NewService("Pechkin", *listenAddr, mem)
	Srv.GetParameters(telegram.ParamTgToken)

	go Srv.Listen()

	b, err := telegram.NewBot(Srv)
	if err.Code != error.ErrCodeNone {
		Srv.Log.Errorf("Error start bot: %s", err)
	}
	b.Listen()

}
