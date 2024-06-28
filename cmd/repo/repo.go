package repo

import (
	"github.com/spf13/cobra"
)

var RepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "manage repos",
	Long:  `manage repos managed by gwt.`,
}

func init() {
	RepoCmd.AddCommand(ListCmd)
	RepoCmd.AddCommand(AddCmd)
	RepoCmd.AddCommand(RemoveCmd)
	RepoCmd.AddCommand(EditCmd)
}
