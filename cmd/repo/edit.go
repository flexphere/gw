package repo

import (
	"cmp"
	"log"
	"os"

	"github.com/flexphere/gw/command"
	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "open config file",
	Long:  `opens the config file with $EDITOR or $VISUAL.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		editor := cmp.Or(os.Getenv("EDITOR"), os.Getenv("VISUAL"))
		if editor == "" {
			log.Fatalln("EDITOR nor VISUAL was set")
		}

		if err := command.PassThrough([]string{editor, configPath}); err != nil {
			log.Fatalf("failed to edit config: %v", err)
		}
	},
}
