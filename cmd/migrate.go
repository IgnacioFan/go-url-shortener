package cmd

import (
	"fmt"
	"go-url-shortener/config"
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
	config, err := config.New()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	db := postgres.InitPostgres(config)
	if err = db.NewMigrate(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
