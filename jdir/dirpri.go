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

import "io/ioutil"

func (d *DirUtils) _GetAllFiles(path string) int {
	iFile := 0

	fReader, err := ioutil.ReadDir(path)

	if err != nil {
		return 0
	}

	for _, fi := range fReader {
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

	fReader, err := ioutil.ReadDir(path)

	if err != nil {
		return 0
	}

	for _, fi := range fReader {
		if !fi.IsDir() {
			iFile++
			d.Files = append(d.Files, path+"/"+fi.Name())

		}
	}

	return iFile
}
