package cmd

import (
	"fmt"
	"go-url-shortener/internal/wire_inject/app"
	"go-url-shortener/pkg/postgres"
	grpccontroller "go-url-shortener/url_shortner_service/controller"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	port        int
	shortUrlCmd = &cobra.Command{
		Use:   "surl",
		Short: "run URL shortener service",
		Run: func(cmd *cobra.Command, args []string) {
			grpccontroller.NewServer()
		},
	}
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
	rootCmd.AddCommand(shortUrlCmd, serverCmd)
	shortUrlCmd.PersistentFlags().IntVarP(&port, "port", "p", 50051, "expose port number")
	serverCmd.PersistentFlags().IntVarP(&port, "port", "p", 3000, "expose port number")
}
