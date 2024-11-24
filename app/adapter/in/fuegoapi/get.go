package fuegoapi

import (
	"archetype/app/shared/infrastructure/fuego/httpserver"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func init() {
	ioc.Registry(newTemplateGet, httpserver.New)
}
func newTemplateGet(s httpserver.Server) {
	fuego.Get(s.Fuego, "/insert-your-custom-pattern-here", func(c *fuego.ContextWithBody[any]) (any, error) {
		body, err := c.Body()
		if err != nil {
			return "unimplemented", err
		}
		return body, nil
	}, option.Summary("newTemplateGet"))
}
