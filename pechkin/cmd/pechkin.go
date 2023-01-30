package main

import (
	"flag"
	"time"

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

	Srv := service.NewService("Pechkin", *listenAddr, mem)
	Srv.GetParameters(telegram.ParamTgToken)

	go Srv.Listen()
	time.Sleep(5 *time.Second)
	b, err := telegram.NewBot(Srv)
	if err.Code != error.ErrCodeNone {
		Srv.Log.Errorf("Error start bot: %s", err)
	}
	b.Listen()

}
