package main

import (
	"github.com/labstack/echo"
	"github.com/ChichaSasha/testProj/api"
)

func main() {
	e := echo.New()

	m := api.NewManager()
	api.Assemble(e, m)

	e.Logger.Fatal(e.Start(":8765"))
}
