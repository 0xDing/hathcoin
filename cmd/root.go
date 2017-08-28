package cmd

import (
	"fmt"
	"os"

	"github.com/borisding1994/hathcoin/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "hathcoin",
	Short: "HathCoin is an experimental digital currency. Long live the man who changed china. Θ..Θ +1s",
	Long: `
那么人呐就都不知道，自己就不可以预料。
你一个人的命运啊，当然要靠自我奋斗，但是也要考虑到历史的行程。

HathCoin is an experimental digital currency.
Long live the man who changed china.`,
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
	config.File = cfgFile
	cobra.OnInitialize(config.InitConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c",
		"", "config file (default is ./config/hathcoin.toml)")

}
