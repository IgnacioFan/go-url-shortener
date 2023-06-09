package cmd

import (
	"fmt"
	"go-url-shortener/internal/wire_inject/app"
	"go-url-shortener/pkg/postgres"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	port      int
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "run HTTP server",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			db, err := postgres.NewPostgres()
			if err = db.NewMigrate(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			app, err := app.Initialize()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			err = app.Start(port)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().IntVarP(&port, "port", "p", 3000, "expose port number")
}
