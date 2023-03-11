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
	// cobra.OnInitialize(initConfig)
	// persistent flags will be global
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "deployment/config/default.ymal", "config file")
	// rootCmd.PersistentFlags().StringP("author", "a", "Weilong Fan", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	// local flags will only run when the action is called
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "Weilong Fan <fan01856472@gmail.com>")
	// viper.SetDefault("license", "apache")
	rootCmd.AddCommand(ServerCmd)
}

// reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".cobra")
// 	}
// 	viper.AutomaticEnv() // read in environment variables that match

// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
