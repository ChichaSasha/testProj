package main

import(
	"github.com/ChichaSasha/testProj/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	m := api.NewManager()
	api.Assemble(e, m)

	e.Logger.Fatal(e.Start(":8765"))
}
