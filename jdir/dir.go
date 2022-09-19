package jdir

import (
	"os"
)

///check file exists
func FileExists(file string) bool {
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
