package httpserver

import (
	"net"
	"net/http"
)

type Server struct {
}

func New() {
}

func (s *Server) POST(string, func(w http.ResponseWriter, r *http.Request)) {
}

func (s *Server) SetListenner(net.Listener) {
}
