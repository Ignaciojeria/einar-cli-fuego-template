package fuegoapi

import (
	"archetype/app/shared/infrastructure/httpserver"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func init() {
	ioc.Registry(newTemplatePut, httpserver.New[*fuego.Server])
}
func newTemplatePut(s httpserver.Server[*fuego.Server]) {
	fuego.Put(s.Manager, "/insert-your-custom-pattern-here", func(c *fuego.ContextWithBody[any]) (any, error) {
		body, err := c.Body()
		if err != nil {
			return "unimplemented", err
		}
		return body, nil
	}, option.Summary("newTemplatePut"))
}
