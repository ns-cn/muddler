package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var CmdVersion = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "打印当前版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("muddler version: v%s(%s/%s)", VERSION, runtime.GOOS, runtime.GOARCH))
	},
}

func RegCmdVersion() {
	root.AddCommand(CmdVersion)
}
