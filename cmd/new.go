package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/flexphere/gw/command"
	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create worktree",
	Long:  `create a new worktree`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()
		conf := config.New(configPath)
		cwd := config.GetCWD()

		if len(args) < 1 || args[0] == "" {
			fmt.Fprintln(os.Stderr, "worktree name is required")
			os.Exit(1)
		}

		repo := conf.FindRepoByPath(cwd)
		if repo == nil {
			fmt.Fprintln(os.Stderr, "not initialized.\nrun `gw init` to initialize repository")
			os.Exit(1)
		}

		//create worktree
		worktreeName := args[0]
		createPath := filepath.Join(repo.WorkDir(), worktreeName)
		cmdArgs := []string{"git", "worktree", "add", "-b", worktreeName, createPath}
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
