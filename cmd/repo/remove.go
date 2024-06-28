/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package repo

import (
	"log"

	"github.com/flexphere/gw/config"

	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove repo",
	Long:  `removes a repo from being managed by gwt.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("please provide a repo name")
		}
		repoName := args[0]
		configPath := config.GetConfigPath()
		conf := config.New(configPath)
		if err := conf.RemoveRepo(repoName); err != nil {
			log.Fatalf("failed to remove repo: %v", err)
		}
	},
}
