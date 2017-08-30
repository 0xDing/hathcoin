package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/borisding1994/hathcoin/utils"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(licenseCmd)
}

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Display license information",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := ioutil.ReadFile("./LICENSE")
		if err != nil {
			utils.Logger.Error(err)
		}
		fmt.Println(string(file))
	},
}
