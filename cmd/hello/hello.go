package hello

import (
	"github.com/SaimonWoidig/golang-training/cmd/root"
	"github.com/SaimonWoidig/golang-training/pkg/hello"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a greeting",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		hello.PrintSayHello(FlagName)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to say hello to",
	)
}
