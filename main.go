package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/iyaozhen/samber-do-learn/service/airplane"
	"github.com/iyaozhen/samber-do-learn/service/car"
	"github.com/samber/do/v2"
)

func main() {
	// create DI container and inject package services
	injector := do.New()

	car.Package(injector)
	airplane.Package(injector)
	//do.ProvideValue(&Config{
	//	Port: 4242,
	//})

	// invoking car will instantiate Car services and its Engine dependency
	car, err := do.Invoke[*car.Car](injector)
	if err != nil {
		log.Fatal(err.Error())
	}

	car.Start() // that's all folk 🤗
	explainInjector := do.ExplainInjector(injector)
	fmt.Println(explainInjector.String())

	// handle ctrl-c and shutdown services
	injector.ShutdownOnSignals(syscall.SIGTERM, os.Interrupt)
}
