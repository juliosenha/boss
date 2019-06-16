package cmd

import (
	"github.com/hashload/boss/core"
	"github.com/spf13/cobra"
)

var dependenciesCmd = &cobra.Command{
	Use:     "dependencies",
	Aliases: []string{"dep"},
	Short:   "Print dependencies tree",
	Long:    `Print dependencies tree and versions`,
	Run: func(cmd *cobra.Command, args []string) {
		core.PrintDependencies()
	},
}

func init() {
	RootCmd.AddCommand(dependenciesCmd)
}