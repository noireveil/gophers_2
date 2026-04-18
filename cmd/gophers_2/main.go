package main

import (
	"gophers_2/internal/delivery/cli"
	"gophers_2/internal/repository"
	"gophers_2/internal/usecase"
)

func main() {
	repo := repository.NewInventoryRepository()
	uc := usecase.NewInventoryUsecase(repo)
	handler := cli.NewHandler(uc)

	handler.Run()
}
