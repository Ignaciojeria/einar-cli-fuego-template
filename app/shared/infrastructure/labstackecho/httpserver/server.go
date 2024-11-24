package httpserver

import (
	"archetype/app/shared/configuration"
	"archetype/app/shared/validator"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
	conf configuration.Conf
}

func init() {
	ioc.Registry(echo.New)
	ioc.Registry(
		New,
		echo.New,
		configuration.NewConf,
		validator.NewValidator)
}

func New(
	e *echo.Echo,
	c configuration.Conf,
	validator *validator.Validator) Server {
	e.Validator = validator
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, time.Second*2)
		defer shutdownCancel()
		if err := e.Shutdown(shutdownCtx); err != nil {
			fmt.Println("Failed to shutdown:", err)
		}
		cancel()
	}()
	return Server{
		conf: c,
		Echo: e,
	}
}

func init() {
	ioc.RegistryAtEnd(Start, New)
}
func Start(e Server) error {
	return e.start()
}

func (s Server) start() error {
	s.printRoutes()
	err := s.Echo.Start(":" + s.conf.PORT)
	fmt.Println(err)
	fmt.Println("waiting for resources to shut down....")
	time.Sleep(2 * time.Second)
	fmt.Println("done.")
	return err
}

func (s Server) printRoutes() {
	routes := s.Echo.Routes()
	for _, route := range routes {
		log.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
}
