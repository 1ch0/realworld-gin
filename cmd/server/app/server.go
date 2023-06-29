package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/1ch0/gin-template/pkg/server/utils/log"

	"github.com/1ch0/gin-template/pkg/server"
	"github.com/1ch0/gin-template/pkg/utils"
	"github.com/spf13/cobra"

	"github.com/1ch0/gin-template/cmd/server/app/options"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:  "run",
		Long: `The API Server services REST operations `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(s)
		},
		SilenceUsage: true,
	}

	return cmd
}

// Run runs the specified APIServer. This should never exit.
func Run(s *options.ServerRunOptions) error {

	errChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if s.GenericServerRunOptions.Server.PprofAddr != "" {
		go utils.EnablePprof(s.GenericServerRunOptions.Server.PprofAddr, errChan)
	}

	go func() {
		if err := run(ctx, s, errChan); err != nil {
			errChan <- fmt.Errorf("failed to run apiserver: %w", err)
		}
	}()
	var term = make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case <-term:
		log.Logger.Infof("Received SIGTERM, exiting gracefully...")
	case err := <-errChan:
		log.Logger.Errorf("Received an error: %s, exiting gracefully...", err.Error())
		return err
	}
	log.Logger.Infof("See you next time!")
	return nil
}

func run(ctx context.Context, s *options.ServerRunOptions, errChan chan error) error {
	app := server.New(*s.GenericServerRunOptions)
	return app.Run()
}
