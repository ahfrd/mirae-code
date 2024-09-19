package cmd

import (
	"context"
	"errors"
	"log"
	"mirae-code/env"
	"mirae-code/infra/network"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type CLI interface {
	Start()
	Error() error
}

type Command struct {
	di   *env.Dependency
	args []string
	err  error
}

func NewCLI(params *env.Dependency, args []string) *Command {
	return &Command{params, args, nil}
}

func (cmd *Command) Start() {

	cmd.runApiGin()

}

func (cmd *Command) runApiGin() {
	serveRoutes := &http.Server{
		Addr:           ":" + cmd.di.Params.Ports.Gin,
		Handler:        network.InitRoutesGin(cmd.di),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if cmd.err = serveRoutes.ListenAndServe(); cmd.err != nil && errors.Is(cmd.err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", cmd.err)
			return
		}

	}()
	cmd.gracefulShutdown(serveRoutes)
	return
}

func (cmd *Command) gracefulShutdown(serverRoutes *http.Server) {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if cmd.err = serverRoutes.Shutdown(ctx); cmd.err != nil {
		return
	}
}

func (cmd *Command) Error() error {
	return cmd.err
}
