package api

import (
	"archetype/app/shared/infrastructure/labstackecho/httpserver"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplatePost, httpserver.New)
}
func newTemplatePost(e httpserver.Server) {
	e.POST("/insert-your-custom-pattern-here", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Unimplemented",
		})
	})
}
