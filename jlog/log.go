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
package jlog

import (
	"os"
)

// LogUtils
// add by yjh 211123
type LogUtils struct {
	_logFile   *os.File
	_logPath   string
	_logName   string
	LastError  error
	_fullPath  string
	_monthLog  bool
	_dayLog    bool
	_singleLog bool
	_logTail   string
}

func BuildDayLogUtils(path string, name string, tail string) *LogUtils {
	return _buildLogUtils(path, name, tail, true, false, false)
}
func BuildMonthLogUtils(path string, name string, tail string) *LogUtils {
	return _buildLogUtils(path, name, tail, false, true, false)
}

func BuildSingleLogUtils(path string, name string, tail string) *LogUtils {
	return _buildLogUtils(path, name, tail, false, false, true)
}

func (l *LogUtils) Close() {
	l._logFile.Close()
}

func (l *LogUtils) Debug(content string) {
	l._appendContent("DEBUG", content)

}
func (l *LogUtils) Info(content string) {
	l._appendContent("INFO", content)

}
func (l *LogUtils) Warn(content string) {
	l._appendContent("WARN", content)

}
func (l *LogUtils) Error(content string) {
	l._appendContent("ERR", content)
}
func (l *LogUtils) Trace(content string) {
	l._appendContent("TRACE", content)
}
