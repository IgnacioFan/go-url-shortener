package cmd

import (
	"fmt"
	"go-url-shortener/internal/adpater/zookeeper"
	"go-url-shortener/internal/app/rest"
	"go-url-shortener/internal/repository/url_repo"
	"go-url-shortener/internal/service/url"
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
      zkClient, err := zookeeper.InitZooKeeper()
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }
      db, err := postgres.NewPostgres()
      if err != nil {
        log.Fatal(err)
        os.Exit(1)
      }
      urlRepo := url_repo.NewShortUrlRepo(db)
    
      urlService, err := url.InitUrl(zkClient, urlRepo)
      if err != nil {
        log.Fatal(err)
        os.Exit(1)
      }
      app := rest.InitShortUrl(urlService)
      if err := app.Run(fmt.Sprintf(":%d", port)); err != nil {
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
