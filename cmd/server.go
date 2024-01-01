package cmd

import (
	"fmt"
	"go-url-shortener/internal/adpater/postgres"
	"go-url-shortener/internal/app/rest"
	"go-url-shortener/internal/service/url_service"
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
      service, err := url_service.NewUrlService()   
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }
      if err := rest.NewRestAPI(port, service); err != nil {
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
