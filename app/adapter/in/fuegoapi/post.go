package fuegoapi

import (
	"archetype/app/shared/infrastructure/httpserver"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func init() {
	ioc.Registry(newTemplatePost, httpserver.New[*fuego.Server])
}
func newTemplatePost(s httpserver.Server[fuego.Server]) {
	fuego.Post(&s.Manager, "/insert-your-custom-pattern-here", func(c *fuego.ContextWithBody[any]) (any, error) {
		body, err := c.Body()
		if err != nil {
			return "unimplemented", err
		}
		return body, nil
	}, option.Summary("newTemplatePost"))
}
