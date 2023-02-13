package main

import (
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/handler"
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/repository/memory"
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/usecase/notify_group"
)

func main() {
	rep := memory.NewNotifyGroupMemoryStorage()
	uc := notify_group.NewNotifyGroupUseCase(rep)
	handlers := handler.NewHandler(uc)

}
