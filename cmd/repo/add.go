package repo

import (
	"log"

	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var (
	worktreePath string
	script       string
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add repo",
	Long:  `adds current git repo to be managed by gwt.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		cwd := config.GetCWD()
		conf := config.New(configPath)
		if err := conf.AddRepo(cwd, worktreePath, script); err != nil {
			log.Fatalf("failed to add repo: %v", err)
		}
	},
}

func init() {
	AddCmd.Flags().StringVarP(&worktreePath, "worktree-path", "w", "", "worktree path")
	AddCmd.Flags().StringVarP(&script, "init-command", "i", "", "command to run upon worktree creation")
}
