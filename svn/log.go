package svn

import (
	"fmt"
	"github.com/ns-cn/muddler/sys"
	"os"
	"strings"
)

type FileUpdate struct {
	IsDeleted bool
	Path      string
}

func Log(branch string, authors []string, revision string, last int) (fileUpdates map[string]bool, err error) {
	if !sys.IsSvnAvailable() {
		return nil, fmt.Errorf("command svn not available")
	}
	_ = os.Chdir(branch)
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("请使用命令行方式运行！%s", err.Error())
	}
	information := Info(workDir)
	//fmt.Println(information)
	logArgs := make([]string, 0)
	logArgs = append(logArgs, "log", "-v")
	if len(revision) != 0 {
		logArgs = append(logArgs, "-r", revision)
	}
	stdout, stderr, err := sys.ExecCommand("svn", logArgs...)
	if err != nil {
		return nil, fmt.Errorf(stderr.String())
	}
	for {
		bytes, err := stderr.ReadBytes('\n')
		if err != nil {
			break
		}
		fmt.Print((string)(bytes))
	}
	var counter = 0
	fileUpdates = make(map[string]bool, last)
	for {
		bytes, err := stdout.ReadBytes('\n')
		if err != nil {
			break
		}
		line := (string)(bytes)
		if strings.HasPrefix(line, "----------") {
			counter++
			if last != 0 && counter > last {
				break
			}
		}
		isDelete := false
		if strings.HasPrefix(line, "   M ") || strings.HasPrefix(line, "   A ") {
		} else if strings.HasPrefix(line, "   D ") {
			isDelete = true
		} else {
			continue
		}
		file := line[5 : len(line)-1]
		if information.UrlRelative != "" {
			if !strings.HasPrefix(file, information.UrlRelative) {
				continue
			} else if information.UrlRelative == file {
				continue
			} else {
				file = strings.ReplaceAll(file, information.UrlRelative, "")
			}
		}
		if strings.Contains(file, " (from ") {
			file = file[:strings.Index(file, " (from ")]
		} else if strings.Contains(file, " (从 ") {
			file = file[:strings.Index(file, " (从 ")]
		}
		_, found := fileUpdates[file]
		if !found {
			fileUpdates[file] = isDelete
		}
	}
	return fileUpdates, nil
}
