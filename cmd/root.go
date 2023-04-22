package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	userLicense string
	cfgFile     string
	rootCmd     = &cobra.Command{
		Use:   "surl",
		Short: "A URL Shortener",
		Long:  "A high available and scalable URL shortener",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	// persistent flags will be global
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	// rootCmd.PersistentFlags().StringP("author", "a", "Weilong Fan", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	// local flags will only run when the action is called
	rootCmd.AddCommand(ServerCmd)
	rootCmd.AddCommand(MigrateCmd)
}
