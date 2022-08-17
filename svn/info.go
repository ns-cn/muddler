package svn

import (
	"bytes"
	"github.com/ns-cn/muddler/svn/constants"
	"os/exec"
	"strings"
)

type Information struct {
	Path               string
	URL                string
	UrlRelative        string
	Root               string
	Uuid               string
	Revision           string
	LastChangeAuthor   string
	LastChangeRevision string
	LastChangeDate     string
}

/*
路径: .
工作副本根目录: /Users/tangyujun/caacetc/fims/r0
URL: https://172.20.21.10:8088/svn/NNG/omms_2018/src/01Code/trunk/jcloud4-dynflight
Relative URL: ^/omms_2018/src/01Code/trunk/jcloud4-dynflight
版本库根: https://172.20.21.10:8088/svn/NNG
版本库 UUID: 3f833441-da7f-904c-a669-db1f547b6d8b
版本: 37542
节点种类: 目录
调度: 正常
最后修改的作者: luomujian
最后修改的版本: 34054
最后修改的时间: 2022-05-31 13:39:31 +0800 (二, 2022-05-31)
*/
func Info(dir string) Information {
	svnInfo := exec.Command("svn", "info")
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	svnInfo.Stdout = &outBuffer
	svnInfo.Stderr = &errBuffer
	err := svnInfo.Run()
	var urlRelative = ""
	if err == nil {
		for {
			readBytes, readErr := outBuffer.ReadBytes('\n')
			if readErr != nil {
				break
			}
			line := (string)(readBytes)
			if strings.HasPrefix(line, constants.H_RELATIVE_URL) {
				urlRelative = line[len(constants.H_RELATIVE_URL) : len(line)-1]
				break
			}
		}
	}
	return Information{UrlRelative: urlRelative}
}
