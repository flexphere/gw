package tree

import (
	"log"

	"github.com/flexphere/gw/command"

	"github.com/spf13/cobra"
)

var TreeCmd = &cobra.Command{
	Use:   "tree",
	Short: "run git worktree commands",
	Long:  `run git worktree commands. any arguments passed to this command will be redirected to the "git worktree" command. For more information on git worktree, see https://git-scm.com/docs/git-worktree`,
	Run: func(cmd *cobra.Command, args []string) {
		git := []string{"git", "worktree"}
		git = append(git, args...)
		err := command.PassThrough(git)
		if err != nil {
			log.Fatal(err)
		}
	},
}
