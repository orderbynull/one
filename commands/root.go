package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const configType = "yaml"
const rootCmdUse = "one"
const flagName = "name"
const flagConfig = "config"
const flagNameUsage = "custom lock name"
const flagConfigUsage = "config file (default is ./one.yaml)"
const rootCmdConfigFailure = "Can't read config:"

// configFile holds path to custom config file.
var configFile string

// RootCmd prints help information
var RootCmd = &cobra.Command{
	Use: rootCmdUse,
}

// initConfig loads config file either from current dir or from --config flag path.
var initConfig = func() {
	viper.SetConfigType(configType)
	viper.SetConfigName(rootCmdUse)

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(rootCmdConfigFailure, err)
		os.Exit(1)
	}
}

// init ...
func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String(flagName, "", flagNameUsage)
	RootCmd.PersistentFlags().StringVar(&configFile, flagConfig, "", flagConfigUsage)
}
