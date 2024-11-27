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
	fuego.Get(s.Manager, "/insert-your-custom-pattern-here",
		func(c *fuego.ContextNoBody) (any, error) {

			return "unimplemented", nil
		}, option.Summary("newTemplateGet"))
}
