package httpserver

import (
	"net"
	"net/http"
)

type Server struct {
}

func New() {
}

func WrapPostStd(Server, string, func(w http.ResponseWriter, r *http.Request)) {
}

func (s *Server) SetListenner(net.Listener) {
}
