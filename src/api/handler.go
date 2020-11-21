package api

import (
	//"fmt"

	"github.com/labstack/echo"
	"net/http"
)

func Assemble(e *echo.Echo, m Manager) {
	h := &handler{
		manager: m,
	}
	e.GET("/:key/status", h.status)
}

type handler struct {
	manager Manager
}

func (h *handler) status(c echo.Context) error {
	//ctx := c.Request().Context()

	key := c.Param("key")
	if key == "" {
		return nil //fmt.Errorf("key is empty")
	}

	result, err := h.manager.GetShortURLStatus(key)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"key": result.Key,
	})
}
