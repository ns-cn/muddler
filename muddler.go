package main

import (
	"github.com/spf13/cobra"
)

const (
	VERSION = "1.0.0"
)

var (
	branch   string   // 仓库位置，不指定则为当前目录
	authors  []string // 过滤用：作者
	revision string   // 修订版本区间，svn方式，可指定具体版本或版本区间
	last     int      // 最近的修订版本次数指定，只查看最近的多少次提交
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
	RegCmdMerge()
	RegCmdVersion()
	_ = root.Execute()
}
