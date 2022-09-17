/*
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
*/
package jutils

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitee.com/yjhi/go/jerrors"
)

/*
*
* add by yjh 211124
* check fiel esists
 */
func FileIsExists(filepath string) bool {
	_, err := os.Lstat(filepath)
	return !os.IsNotExist(err)
}

/*
*
* add by yjh 211124
* read file content
 */
func ReadFileAll(path string) (string, error) {

	if !FileIsExists(path) {
		return "", jerrors.NewError("文件不存在")
	}

	f, err := os.Open(path)
	if err != nil {
		s := fmt.Sprintf("读取文件失败:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		s := fmt.Sprintf("获取文件内容失败:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	return string(fd), nil
}

/*
*
* add by yjh 211124
* make dir
 */
func MkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
		os.Chmod(path, 0777)
	}
}
