package main

import (
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/handler"
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/repository/memory"
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/usecase/notify_group"
	service_V1 "github.com/Alang0r/vypolnyator/pkg/service_v1"
)

func main() {
	rep := memory.NewNotifyGroupMemoryStorage()
	uc := notify_group.NewNotifyGroupUseCase(rep)
	handlers:= handler.NewHandler(uc)

	srv := service_V1.NewService("skald", ":3003")
	srv.AddHAndlers(handlers)
	srv.Listen()
}
