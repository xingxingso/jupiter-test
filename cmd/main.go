package main

import (
	"fmt"
	"github.com/douyu/jupiter"
	"jupiter-test/internal/app/engine"
	"jupiter-test/internal/app/model"
	"jupiter-test/internal/app/service"
	"log"
)

func main() {
	eng := engine.NewEngine()
	eng.RegisterHooks(jupiter.StageAfterStop, func() error {
		fmt.Println("exit jupiter app ...")
		return nil
	})

	model.Init()
	service.Init()
	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}
