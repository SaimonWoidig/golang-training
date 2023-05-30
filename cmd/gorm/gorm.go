package gorm

import (
	"fmt"

	"github.com/SaimonWoidig/golang-training/cmd/root"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "gorm",
	Short: "GORM example",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		gormFunc()
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}

type Pet struct {
	gorm.Model
	Name string
	Age  int
}

func gormFunc() {
	log.Info().Msg("Starting GORM")
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Pet{})
	db.Create(&Pet{Name: "Bob", Age: 1})

	var pet Pet
	db.First(&pet, 1)
	fmt.Printf("%s %d\n", pet.Name, pet.Age)

	db.Create(&Pet{Name: "Steve", Age: 5})
	db.Create(&Pet{Name: "Alice", Age: 2})

	var pets []Pet
	db.Find(&pets)
	for i, p := range pets {
		fmt.Printf("%d %s %d\n", i, p.Name, p.Age)
	}
}
