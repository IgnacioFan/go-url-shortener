package cmd

import (
	"fmt"
	"go-url-shortener/internal/delivery"

	"github.com/spf13/cobra"
)

var (
	ServerCmd = &cobra.Command{
		Use:           "server",
		Short:         "Initialize Url shortener",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run:           runServerCmd,
	}
)

func runServerCmd(cmd *cobra.Command, args []string) {
	application := delivery.NewHttpServer()
	application.Run(fmt.Sprintf(":%d", 3000))
}
