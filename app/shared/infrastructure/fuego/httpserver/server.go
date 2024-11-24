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
	ioc.RegistryAtEnd(startAtEnd, New)
}

type Server struct {
	Manager *fuego.Server
	conf    configuration.Conf
}

func New(conf configuration.Conf) Server {
	server := Server{
		Manager: fuego.NewServer(fuego.WithAddr(":" + conf.PORT)),
		conf:    conf,
	}
	server.healthCheck()
	return server
}

func startAtEnd(e Server) error {
	return e.Manager.Run()
}

func (s Server) healthCheck() error {
	h, err := health.New(
		health.WithComponent(health.Component{
			Name:    s.conf.PROJECT_NAME,
			Version: s.conf.VERSION,
		}), health.WithSystemInfo())
	if err != nil {
		return err
	}
	fuego.GetStd(s.Manager,
		"/health",
		h.Handler().ServeHTTP,
		option.Summary("healthCheck"))
	return nil
}

func WrapPostStd(s Server, path string, f func(w http.ResponseWriter, r *http.Request)) {
	fuego.PostStd(s.Manager, path, f)
}
