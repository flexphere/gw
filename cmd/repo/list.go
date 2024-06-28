package repo

import (
	"fmt"

	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "prints list of repos",
	Long:  `prints list of repos managed by gwt.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		conf := config.New(configPath)
		for _, value := range conf.Config {
			fmt.Println(value.Name())
		}
	},
}
