package cmd

import (
	"github.com/SaimonWoidig/golang-training/cmd/root"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
