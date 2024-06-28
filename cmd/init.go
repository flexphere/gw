package cmd

import (
	"fmt"
	"os"

	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var (
	worktreePath string
	script       []string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize current repository to be managed by gw",
	Long:  "initialize current repository to be managed by gw",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		cwd := config.GetCWD()
		conf := config.New(configPath)
		commands, err := cmd.Flags().GetStringArray("cmd")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := conf.AddRepo(cwd, worktreePath, commands); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	initCmd.Flags().StringVarP(&worktreePath, "worktree-path", "w", "", "worktree path(default: ./.worktrees)")
	initCmd.Flags().StringArrayVar(&script, "cmd", []string{}, "command to run upon worktree creation")
}
