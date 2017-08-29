package cmd

import (
	"fmt"
	"net"

	"github.com/borisding1994/hathcoin/config"
	"github.com/borisding1994/hathcoin/utils"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start HathCoin Server",
	RunE: func(cmd *cobra.Command, arg []string) error {
		addr := config.GetString("addr")
		l, err := net.Listen("tcp", addr)
		if err != nil {
			utils.Logger.Fatal(err)
			return err
		}
		s := grpc.NewServer()
		// Register reflection service on gRPC server.
		//
		// gRPC Server Reflection provides information about publicly-accessible gRPC services on a server,
		// and assists clients at runtime to construct RPC requests and responses without precompiled service information.
		reflection.Register(s)
		// [maybe strange bug]: `fmt.Printf` and `Logger` not work on `s.Serve(l)` after.
		utils.Logger.Infof("HathCoin Server is work. bind address is: %s", addr)
		fmt.Printf("\n\nPress Ctrl + C to stop server.\n\n")
		if err := s.Serve(l); err != nil {
			utils.Logger.Fatal(err)
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
