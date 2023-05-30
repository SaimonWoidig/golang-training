package server

import (
	"fmt"

	"github.com/SaimonWoidig/golang-training/cmd/root"
	"github.com/SaimonWoidig/golang-training/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Simple HTTP server",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		addr := fmt.Sprintf(":%d", FlagPort)
		server.Server(addr)
	},
}

func init() {
	viper.BindEnv("PORT")
	port := viper.GetInt("PORT")

	root.Cmd.AddCommand(Cmd)
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
}
