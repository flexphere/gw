package cmd

import (
	"cmp"
	"fmt"
	"os"

	"github.com/flexphere/gw/command"
	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "open config file",
	Long:  `opens the config file with $EDITOR or $VISUAL.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		editor := cmp.Or(os.Getenv("EDITOR"), os.Getenv("VISUAL"))
		if editor == "" {
			fmt.Fprintln(os.Stderr, "EDITOR nor VISUAL was set")
			os.Exit(1)
		}

		if err := command.PassThrough([]string{editor, configPath}); err != nil {
			fmt.Fprintln(os.Stderr, "failed to edit config")
			os.Exit(1)
		}
	},
}
