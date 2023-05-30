package version

import (
	"fmt"

	"github.com/SaimonWoidig/golang-training/cmd/root"
	"github.com/SaimonWoidig/golang-training/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Prints version",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf(version.Version)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
