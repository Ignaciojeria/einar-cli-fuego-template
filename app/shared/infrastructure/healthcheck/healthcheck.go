package api

import (
	"archetype/app/shared/configuration"
	"archetype/app/shared/infrastructure/labstackecho/httpserver"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/hellofresh/health-go/v5"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(healthCheck,
		httpserver.New,
		configuration.NewConf)
}

// To see usage examples of the library, visit: https://github.com/hellofresh/health-go
func healthCheck(e httpserver.Server, c configuration.Conf) {
	h, _ := health.New(
		health.WithComponent(health.Component{
			Name:    c.PROJECT_NAME,
			Version: c.VERSION,
		}), health.WithSystemInfo())
	e.GET("/health", echo.WrapHandler(h.Handler()))
}