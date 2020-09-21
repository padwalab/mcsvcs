package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"github.com/padwalab/mcsvcs/pkg/edil"
	"github.com/padwalab/mcsvcs/pkg/edil/endpoints"
	"github.com/padwalab/mcsvcs/pkg/edil/transport"
)

const defaultHTTPPort = "8081"

func main() {
	var (
		httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT_WATERMARK", defaultHTTPPort))
	)

	var (
		service     = edil.NewService()
		eps         = endpoints.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(eps)
	)
	var g group.Group
	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			fmt.Println("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			fmt.Println("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	fmt.Println(g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv("env")
	if e == "" {
		return fallback
	}
	return e
}
