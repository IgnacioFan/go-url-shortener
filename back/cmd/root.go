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
		Short: "URL Shortener",
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
}
