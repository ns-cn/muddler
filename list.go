package main

import (
	"fmt"
	"github.com/ns-cn/muddler/svn"
	"github.com/spf13/cobra"
	"os"
)

var (
	branch   string   // 仓库位置，不指定则为当前目录
	authors  []string // 过滤用：作者
	revision string   // 修订版本区间，svn方式，可指定具体版本或版本区间
	last     int      // 最近的修订版本次数指定，只查看最近的多少次提交
)
var cmdList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "查询变更",
	Long:    `查询svn提交记录的文件变更`,
	Run: func(cmd *cobra.Command, args []string) {
		fileUpdates, err := svn.Log(branch, authors, revision, last)
		if err != nil {
			_, _ = os.Stderr.Write([]byte(err.Error()))
			cmd.Help()
			return
		}
		for _, fileUpdate := range fileUpdates {
			if fileUpdate.IsDeleted {
				fmt.Println("deleted\t" + fileUpdate.Path)
			} else {
				fmt.Println("updated\t" + fileUpdate.Path)
			}
		}
	},
}

func RegCmdList() {
	cmdList.Flags().StringSliceVarP(&authors, "authors", "a", []string{}, "filter by authors")
	cmdList.Flags().StringVarP(&branch, "branch", "b", ".", "source branch")
	cmdList.Flags().StringVarP(&revision, "revision", "r", "", "版本，不指定则为所有，可选单次(1024)或范围(1024:2048)")
	cmdList.Flags().IntVarP(&last, "last", "l", 0, "指定查看最近的多少次更新")
	root.AddCommand(cmdList)
}
