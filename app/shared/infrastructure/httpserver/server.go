package httpserver

import (
	"archetype/app/shared/configuration"
	"net"
	"net/http"
)

type Server struct {
	Manager any
}

func New(configuration.Conf) Server {
	return Server{}
}

func WrapPostStd(Server, string, func(w http.ResponseWriter, r *http.Request)) {
}

func (s *Server) SetListenner(net.Listener) {
}
