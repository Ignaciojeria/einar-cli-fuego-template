package api

import (
	"archetype/app/shared/infrastructure/httpserver"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplateDelete, httpserver.New[*echo.Echo])
}
func newTemplateDelete(e httpserver.Server[*echo.Echo]) {
	e.Manager.DELETE("/insert-your-custom-pattern-here", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Unimplemented",
		})
	})
}
