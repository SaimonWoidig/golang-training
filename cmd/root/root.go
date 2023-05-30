package root

import (
	"github.com/SaimonWoidig/golang-training/version"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "golang-training",
	Short: "golang-training " + version.Version,
}
