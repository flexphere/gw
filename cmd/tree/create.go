package tree

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/flexphere/gw/command"
	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create worktree",
	Long:  `create a new worktree`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		conf := config.New(configPath)
		cwd := config.GetCWD()

		if len(args) < 1 || args[0] == "" {
			log.Fatalf("empty worktree name")
		}

		repo := conf.FindRepoByPath(cwd)
		if repo == nil {
			log.Fatalf("not in a managed repo")
		}

		//create worktree
		worktreeName := args[0]
		createPath := filepath.Join(repo.WorkDir(), worktreeName)
		cmdArgs := []string{"git", "worktree", "add", "--orphan", "-b", worktreeName, createPath}
		if err := command.PassThrough(cmdArgs); err != nil {
			log.Fatalf("failed to run `%s`: %v", strings.Join(cmdArgs, " "), err)
		}

		//run init command
		if len(repo.Cmd()) > 0 {
			for _, c := range repo.Cmd() {
				if c == "" {
					continue
				}
				cmdArgs := []string{"sh", "-c", c}
				if err := command.PassThrough(cmdArgs); err != nil {
					log.Fatalf("failed to run `%s`: %v", strings.Join(cmdArgs, " "), err)
				}
			}
		}

		fmt.Println(createPath)
	},
}

func init() {
	TreeCmd.AddCommand(CreateCmd)
}
