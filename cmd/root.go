package cmd

import (
	"os"

	"github.com/flexphere/gw/cmd/repo"
	"github.com/flexphere/gw/cmd/tree"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gw",
	Short: "gw is a git worktree command wrapper",
	Long:  `gw is a git worktree command wrapper`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(repo.RepoCmd)
	rootCmd.AddCommand(tree.TreeCmd)
}
