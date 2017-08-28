package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// variables set at build time by Makefile
var (
	CommitHash string // CommitHash = `git rev-parse --short HEAD 2>/dev/null || echo "unreleased" `
	Version    string // Version = `cat $(CURDIR)/VERSION`
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of HathCoin",
	Long:  `All software has versions. This is HathCoin's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("HathCoin current version is \033[32m%s-%s\033[0m\n", Version, CommitHash)
	},
}
