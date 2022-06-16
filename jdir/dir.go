package dir

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func FileExist(file string) bool {
	_, err := os.Lstat(file)

	return !os.IsNotExist(err)

}

type DirUtils struct {
	Files []string
	Path  string
}

func BuildDirUtils(path string) *DirUtils {
	return &DirUtils{
		Files: make([]string, 20),
		Path:  path,
	}
}

func (d *DirUtils) _GetAllFiles(path string) int {
	iFile := 0

	freader, err := ioutil.ReadDir(path)

	if err != nil {
		return 0
	}

	for _, fi := range freader {
		if !fi.IsDir() {
			iFile++
			d.Files = append(d.Files, path+"/"+fi.Name())
		} else {
			iFile += d._GetAllFiles(path + "/" + fi.Name())
		}
	}

	return iFile

}

func (d *DirUtils) _GetDirFiles(path string) int {

	iFile := 0

	freader, err := ioutil.ReadDir(path)

	if err != nil {
		return 0
	}

	for _, fi := range freader {
		if !fi.IsDir() {
			iFile++
			d.Files = append(d.Files, path+"/"+fi.Name())

		}
	}

	return iFile
}

func (d *DirUtils) GetAllFiles() int {
	if len(d.Files) > 0 {
		d.Files = d.Files[0:0]
	}
	return d._GetAllFiles(d.Path)
}

func (d *DirUtils) GetDirFiles() int {
	if len(d.Files) > 0 {
		d.Files = d.Files[0:0]
	}
	return d._GetDirFiles(d.Path)
}

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
