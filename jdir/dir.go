/****************************************************************************
MIT License

Copyright (c) 2022 yjhi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*****************************************************************************/

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
