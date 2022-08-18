# muddler
muddler: merge helper for SVN

核心功能：查看svn提交记录的文件变更、合并文件变更到其他目录
> 用于解决svn命令行合并代码的不便

## 查看svn提交记录的文件变更
```shell
> muddler list -h
功能： 【查询svn提交记录的文件变更】
必要参数：源分支branch(b)、版本范围（last(l)或revision(r)）
查看仓库最近10次提交： muddler list -b ./code/svn -l 10
查看仓库指定版本号的提交： muddler list -b ./code/svn -r 1024
查看仓库指定版本区间的提交： muddler list -b ./code/svn -r 1024:2048
可选增加preview用于是否预览，还是直接合并

Usage:
   list [flags]

Aliases:
  list, l

Flags:
  -a, --authors strings   通过提交者过滤（暂未生效）
  -b, --branch string     源分支 (default ".")
  -h, --help              help for list
  -l, --last int          指定查看最近的多少次更新
  -r, --revision string   版本，不指定则为所有，可选单次(1024)或范围(1024:2048)
```

## 合并svn提交记录的文件变更

```shell
> muddler merge -h
功能： 【查询svn提交记录的文件变更并合并到指定目录】
必要参数：源分支branch(b)、目标分支target(t,可多选)、版本范围（last(l)或revision(r)）
存放仓库最近10次提交到指定目录： muddler list -b ./code/svn -t ./update-20220818 -l 10
合并指定分支版本到另外多个分支： muddler list -b ./code/svn -t ./branch1 -t ./branch2 -r 1024
合并仓库指定版本区间的提交到另外一个分支： muddler list -b ./code/svn -r 1024:2048 -t ./branch1

Usage:
   merge [flags]

Aliases:
  merge, m

Flags:
  -a, --authors strings   通过提交者过滤（暂未生效）
  -b, --branch string     源分支 (default ".")
  -h, --help              help for merge
  -l, --last int          最近的多少次更新
  -p, --preview           是否预览所有的变更，否则直接提交 (default true)
  -r, --revision string   版本，不指定则为所有，可选单次(1024)或范围(1024:2048)
  -t, --target strings    目标分支
```

> 注：Windows 端实测存在各种问题，本人也不想花过多时间在windows上
> 
> 小乌龟自带本命令行的图形界面功能，所以不再针对Windows提供功能