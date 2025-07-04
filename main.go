package main

import (
	"log"

	"ping/api"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

	e := echo.New()

	api.RegisterHandlers(e, server)

	// And we serve HTTP until the world ends.s
	log.Fatal(e.Start("0.0.0.0:8080"))
}
