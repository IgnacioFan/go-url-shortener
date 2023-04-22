package cmd

import (
	"fmt"
	"go-url-shortener/internal/wire_inject/app"
	"os"

	"github.com/spf13/cobra"
)

var (
	ServerCmd = &cobra.Command{
		Use:           "server",
		Short:         "run Url Shortener",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run:           runServerCmd,
	}
)

func runServerCmd(cmd *cobra.Command, args []string) {
	app, err := app.Initialize()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = app.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
