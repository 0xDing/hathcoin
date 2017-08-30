package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// InitConfig reads in config file and ENV variables if set.
func init() {
	if file := os.Getenv("HAC_CONFIG"); file != "" {
		// Use config file from the flag.
		viper.SetConfigFile(file)
	} else {
		// look for config in the directory of the currently running file (without extension).
		viper.AddConfigPath("./config")
		viper.SetConfigName("hathcoin")
	}

	viper.SetEnvPrefix("HAC")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
}

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface.
var Get = viper.Get

// GetInt returns the value associated with the key as an integer.
var GetInt = viper.GetInt

// GetString returns the value associated with the key as a string.
var GetString = viper.GetString

// GetBool returns the value associated with the key as a boolean.
var GetBool = viper.GetBool
