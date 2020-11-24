package main

import (
	"fmt"
	"log"
	"os"

	"github.com/omekov/golang-ci/internal/handlers"
	"github.com/omekov/golang-ci/pkg/healthy"
)

func heathCheck() {
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		if err := healthy.Check(); err != nil {
			log.Fatal(err.Error())
		}
		os.Exit(0)
	}
}

func main() {
	fmt.Print("stating server")
	handlers.Mux()
	heathCheck()
}
