package root

import (
	"github.com/SaimonWoidig/golang-training/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "golang-training",
	Short: "golang-training " + version.Version,
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	viper.SetEnvPrefix("TRAINING")
}
