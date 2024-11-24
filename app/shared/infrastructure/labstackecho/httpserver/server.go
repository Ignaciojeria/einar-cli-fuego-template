package httpserver

import (
	"archetype/app/shared/configuration"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/hellofresh/health-go/v5"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Manager *echo.Echo
	conf    configuration.Conf
}

func init() {
	ioc.Registry(New, configuration.NewConf)
	ioc.Registry(
		healthCheck,
		New,
		configuration.NewConf)
	ioc.RegistryAtEnd(Start, New)
}

func New(c configuration.Conf) Server {
	e := echo.New()
	e.Validator = NewValidator()
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
		conf:    c,
		Manager: e,
	}
}

func Start(s Server) error {
	printRoutes(s)
	err := s.Manager.Start(":" + s.conf.PORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	fmt.Println("waiting for resources to shut down....")
	time.Sleep(2 * time.Second)
	fmt.Println("done.")
	return err
}

func WrapPostStd(s Server, path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Manager.POST(path, echo.WrapHandler(http.HandlerFunc(f)))
}

func printRoutes(s Server) {
	routes := s.Manager.Routes()
	for _, route := range routes {
		log.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
}

// To see usage examples of the library, visit: https://github.com/hellofresh/health-go
func healthCheck(e Server, c configuration.Conf) {
	h, _ := health.New(
		health.WithComponent(health.Component{
			Name:    c.PROJECT_NAME,
			Version: c.VERSION,
		}), health.WithSystemInfo())
	e.Manager.GET("/health", echo.WrapHandler(h.Handler()))
}

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}
