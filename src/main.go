package main

import (
	"github.com/ChichaSasha/testProj/api"

	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	m := api.NewManager()
	api.Assemble(e, m)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
