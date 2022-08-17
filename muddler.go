package main

import (
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Short: "muddler: merge helper for SVN",
	Long: `muddler: merge helper for SVN, 
command svn required for common usage`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	RegCmdList()
	root.Execute()
}
