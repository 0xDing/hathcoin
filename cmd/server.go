package cmd

import (
	"fmt"

	"github.com/borisding1994/hathcoin/core"
	"github.com/borisding1994/hathcoin/utils"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start HathCoin Server",
	RunE: func(cmd *cobra.Command, arg []string) error {
		core.Run()

		utils.Logger.Infof("HathCoin Server is work. ")
		fmt.Printf("\n\nPress Ctrl + C to stop server.\n\n")

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
