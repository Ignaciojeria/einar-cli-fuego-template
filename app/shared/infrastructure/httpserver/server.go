package httpserver

import (
	"archetype/app/shared/configuration"
	"net"
	"net/http"
)

type Server[T any] struct {
	Manager T
}

func New[T any](configuration.Conf) Server[T] {
	return Server[T]{}
}

func WrapPostStd(Server[any], string, func(w http.ResponseWriter, r *http.Request)) {
}

func (s *Server[any]) SetListenner(net.Listener) {
}
