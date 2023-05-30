package server

import (
	"fmt"

	"github.com/SaimonWoidig/golang-training/cmd/root"
	"github.com/SaimonWoidig/golang-training/pkg/ginsrv"
	"github.com/SaimonWoidig/golang-training/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagPort int
var FlagGin bool

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Simple HTTP server",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		address := fmt.Sprintf(":%d", FlagPort)
		if FlagGin {
			ginsrv.Server(address)
		} else {
			server.Server(address)
		}
	},
}

func init() {
	viper.BindEnv("PORT")
	port := viper.GetInt("PORT")

	root.Cmd.AddCommand(Cmd)
	// port flag
	Cmd.Flags().IntVarP(
		&FlagPort,
		"port",
		"p",
		port,
		"On which port to listen to",
	)
	if port == 0 {
		Cmd.MarkFlagRequired("port")
	}

	// gin flag
	Cmd.Flags().BoolVarP(
		&FlagGin,
		"gin",
		"g",
		true,
		"Whether to run server as a Gin server")
}
