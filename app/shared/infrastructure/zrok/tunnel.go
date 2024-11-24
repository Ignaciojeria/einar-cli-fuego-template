package ngrok

import (
	"archetype/app/shared/infrastructure/httpserver"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/openziti/zrok/environment"
	"github.com/openziti/zrok/sdk/golang/sdk"
)

func init() {
	ioc.Registry(NewTunnel, httpserver.New[any])
}

func NewTunnel(s httpserver.Server[any]) error {
	root, err := environment.LoadRoot()
	if err != nil {
		return err
	}
	shr, err := sdk.CreateShare(root, &sdk.ShareRequest{
		BackendMode: sdk.ProxyBackendMode,
		ShareMode:   sdk.PublicShareMode,
		Frontends:   []string{"public"},
		Target:      "http-server",
	})

	if err != nil {
		return err
	}

	listenner, err := sdk.NewListener(shr.Token, root)
	if err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		if err := sdk.DeleteShare(root, shr); err != nil {
			fmt.Println("Failed to zrok tunnel shutdown:", err)
		}
	}()
	s.SetListenner(listenner)
	return nil
}
