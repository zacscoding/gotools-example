package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zacscoding/gotools-example/person"
	personDB "github.com/zacscoding/gotools-example/person/database"
	"log"
)

const (
	serverPort = 8800
)

func main() {
	e := echo.New()

	personHandler, _ := person.NewHandler(personDB.NewInmemoryPersonDB())
	personHandler.Route(e)

	if err := e.Start(fmt.Sprintf(":%d", serverPort)); err != nil {
		log.Fatal(err)
	}
}
