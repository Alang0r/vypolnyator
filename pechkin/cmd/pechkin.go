package main

import (
	"flag"

	_ "github.com/Alang0r/vypolnyator/pechkin/internal/handlers"
	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/Alang0r/vypolnyator/pkg/telegram"
)

var teleHandlers map[string] telegram.THandler

func main() {
	mem := storage.NewMemoryStorage()
	listenAddr := flag.String("listenaddr", ":3002", "listening address")
	flag.Parse()

	srv := service.NewService("Pechkin", *listenAddr, mem)
	srv.GetParameters(telegram.ParamTgToken)

	go srv.Listen()

	b, err := telegram.NewBot(srv)
	if err.Code != error.ErrCodeNone {
		srv.Log.Errorf("Error start bot: %s", err)
	}

	b.Listen()

}
