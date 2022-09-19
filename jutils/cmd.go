package jutils

import (
	"io/ioutil"
	"os/exec"
	"runtime"
	"syscall"
)

func RunCmd(c string, p string) (error, string) {

	cmd := exec.Command(c, p)

	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err, ""
	}

	defer stdout.Close()

	if err := cmd.Start(); err != nil {

		return err, ""

	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {

		return err, ""

	} else {

		return nil, string(opBytes)

	}
}
