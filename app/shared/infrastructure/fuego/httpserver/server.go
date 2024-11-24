package httpserver

import (
	"archetype/app/shared/configuration"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/hellofresh/health-go/v5"
)

func init() {
	ioc.Registry(New, configuration.NewConf)
	ioc.Registry(healthCheck, New, configuration.NewConf)
	ioc.RegistryAtEnd(startAtEnd, New)
}

type Server struct {
	Fuego *fuego.Server
}

func New(conf configuration.Conf) Server {
	return Server{fuego.NewServer(fuego.WithAddr(":" + conf.PORT))}
}

func startAtEnd(e *fuego.Server) error {
	return e.Run()
}

func healthCheck(s *fuego.Server, c configuration.Conf) error {
	h, err := health.New(
		health.WithComponent(health.Component{
			Name:    c.PROJECT_NAME,
			Version: c.VERSION,
		}), health.WithSystemInfo())
	if err != nil {
		return err
	}
	fuego.GetStd(s,
		"/health",
		h.Handler().ServeHTTP,
		option.Summary("healthCheck"))
	return nil
}

func WrapPostStd(s Server, path string, f func(w http.ResponseWriter, r *http.Request)) {
	fuego.PostStd(s.Fuego, path, f)
}
