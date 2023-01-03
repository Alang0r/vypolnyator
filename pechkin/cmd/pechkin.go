package main

import (
	"flag"
	"log"

	service "github.com/Alang0r/vypolnyator/pkg/service"
	storage "github.com/Alang0r/vypolnyator/pkg/storage"
	telegram "github.com/Alang0r/vypolnyator/pkg/telegram"
)

func main() {

	mem := storage.NewMemoryStorage()
	listenAddr := flag.String("listenaddr", ":3001", "listening address")
	flag.Parse()

	srv := service.NewService("Pechkin", *listenAddr, mem)

	srv.Start()
	log.Printf("Pechkin is listening on port: %s", *listenAddr)
	
	params, err := telegram.GetParameters("https://api.telegram.org/bot", "token")
	if err != nil {
		log.Println(err)
	}

	tg, err := telegram.NewTelegramCommunicator(*params)
	if err != nil {
		log.Println(err)
	}

	tg.ListenAndServe()

}
