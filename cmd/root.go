package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "hathcoin",
	Short: "HathCoin is an experimental digital currency, Just for learning blockchain and golang.",
	Long: `
垂死病中惊坐起，谈笑风生又一年

HathCoin is an experimental digital currency, Just for learning blockchain and golang.
"Hath(蛤丝)" is a Chinese Internet meme. Long live the man who changed china.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	os.Setenv("HAC_CONFIG", cfgFile)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c",
		"", "config file (default is ./config/hathcoin.toml)")
}
