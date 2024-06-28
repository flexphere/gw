package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/flexphere/gw/command"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gw",
	Short: "gw is a git worktree command wrapper",
	Long:  `gw is a git worktree command wrapper`,
	Run: func(cmd *cobra.Command, args []string) {
		git := []string{"git", "worktree"}
		git = append(git, args...)
		err := command.PassThrough(git)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(initCmd)
}
