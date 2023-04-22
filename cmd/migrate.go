package cmd

import (
	"fmt"
	"go-url-shortener/pkg/postgres"
	"os"

	"github.com/spf13/cobra"
)

var (
	MigrateCmd = &cobra.Command{
		Use:           "migrate",
		Short:         "Migrate database",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run:           runMigrateCmd,
	}
)

func runMigrateCmd(cmd *cobra.Command, args []string) {
	err := postgres.NewMigrate()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
