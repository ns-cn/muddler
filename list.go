package main

import (
	"fmt"
	"github.com/ns-cn/muddler/svn"
	"github.com/spf13/cobra"
	"os"
)

var cmdList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "查询变更",
	Long: `功能： 【查询svn提交记录的文件变更】
必要参数：源分支branch(b)、版本范围（last(l)或revision(r)）
查看仓库最近10次提交： muddler list -b ./code/svn -l 10
查看仓库指定版本号的提交： muddler list -b ./code/svn -r 1024
查看仓库指定版本区间的提交： muddler list -b ./code/svn -r 1024:2048
可选增加
`,
	Run: func(cmd *cobra.Command, args []string) {
		fileUpdates, err := svn.Log(branch, authors, revision, last)
		if err != nil {
			_, _ = os.Stderr.Write([]byte(err.Error()))
			cmd.Help()
			return
		}
		for file, isDelete := range fileUpdates {
			if isDelete {
				fmt.Println("deleted\t" + file)
			} else {
				fmt.Println("updated\t" + file)
			}
		}
	},
}

func RegCmdList() {
	cmdList.Flags().StringSliceVarP(&authors, "authors", "a", []string{}, "通过提交者过滤（暂未生效）")
	cmdList.Flags().StringVarP(&branch, "branch", "b", ".", "源分支")
	cmdList.Flags().StringVarP(&revision, "revision", "r", "", "版本，不指定则为所有，可选单次(1024)或范围(1024:2048)")
	cmdList.Flags().IntVarP(&last, "last", "l", 0, "指定查看最近的多少次更新")
	root.AddCommand(cmdList)
}
