package api

import (
	"archetype/app/shared/infrastructure/httpserver"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplatePut, httpserver.New[*echo.Echo])
}
func newTemplatePut(e httpserver.Server[*echo.Echo]) {
	e.Manager.PUT("/insert-your-custom-pattern-here", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Unimplemented",
		})
	})
}
