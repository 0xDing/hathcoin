package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// File is configfile
var File string

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if File != "" {
		// Use config file from the flag.
		viper.SetConfigFile(File)
	} else {
		// look for config in the directory of the currently running file (without extension).
		appPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		viper.AddConfigPath(filepath.Join(appPath, "config"))
		viper.SetConfigName("hathcoin")
	}

	viper.SetEnvPrefix("HAC")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}
}

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface.
func Get(key string) interface{} {
	return viper.Get(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
	return viper.GetBool(key)
}
