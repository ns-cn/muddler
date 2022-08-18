package sys

import (
	"bytes"
	"os/exec"
)

func ExecCommand(name string, arg ...string) (bytes.Buffer, bytes.Buffer, error) {
	cmd := exec.Command(name, arg...)
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	return outBuffer, errBuffer, err
}

func IsSvnAvailable() bool {
	svnCheck := exec.Command("svn", "help")
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	svnCheck.Stdout = &outBuffer
	svnCheck.Stderr = &errBuffer
	err := svnCheck.Start()
	return err == nil
}
