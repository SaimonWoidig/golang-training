package cmd

import (
	_ "github.com/SaimonWoidig/golang-training/cmd/hello"
	"github.com/SaimonWoidig/golang-training/cmd/root"
	_ "github.com/SaimonWoidig/golang-training/cmd/server"
	_ "github.com/SaimonWoidig/golang-training/cmd/version"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
