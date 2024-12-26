package fuegoapi

import (
	"archetype/app/shared/infrastructure/fuego/httpserver"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func init() {
	ioc.Registry(newTemplatePatch, httpserver.New)
}
func newTemplatePatch(s httpserver.Server) {
	fuego.Patch(s.Manager, "/insert-your-custom-pattern-here",
		func(c fuego.ContextNoBody) (any, error) {

			return "unimplemented", nil
		}, option.Summary("newTemplatePatch"))
}
