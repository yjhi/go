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
package jlog2

import "gitee.com/yjhi/go/jlog"

var LogRoot string = "logs"
var LogName string = "day"
var LogTail string = "log"
var _logUtil *jlog.LogUtils = nil

func _initLog() {
	if _logUtil == nil {
		_logUtil = jlog.BuildDayLogUtils(LogRoot, LogName, LogTail)
	}
}

func LogDebug(text string) {
	_initLog()

	_logUtil.Debug(text)
}

func LogInfo(text string) {
	_initLog()

	_logUtil.Info(text)
}
func LogError(text string) {
	_initLog()

	_logUtil.Error(text)
}

func LogWarn(text string) {
	_initLog()

	_logUtil.Warn(text)
}