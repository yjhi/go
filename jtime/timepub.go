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

package jtime

import (
	"fmt"
	"time"
)

/*
* add by yjh 211124
* build
 */

func BuildTimeUtils() *TimeUtils {
	return &TimeUtils{
		FormatDate:     "2006-01-02",
		FormatTime:     "15:04:05",
		FormatDateTime: "2006-01-02 15:04:05",
	}
}

/*
* add by yjh 211124
* build
 */
func BuildTimeUtilsNoFmt() *TimeUtils {
	return &TimeUtils{
		FormatDate:     "20060102",
		FormatTime:     "150405",
		FormatDateTime: "20060102150405",
	}
}

func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowTime() string {
	return time.Now().Format("15:04:05")
}

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func UnixTime() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func BuildTimeCount() *TimeCount {
	return &TimeCount{
		_startTime: time.Now(),
	}
}

// parse timestr to time
// add by yjh
// 211202
func ParseDateTime(f string, s string) (time.Time, error) {

	return time.ParseInLocation(_formatDateTimeParse(f), s, time.Local)
}

func ParseTime(f string, s string) (time.Time, error) {

	return time.ParseInLocation(_formatTimeParse(f), s, time.Local)
}

func ParseDate(f string, s string) (time.Time, error) {

	return time.ParseInLocation(_formatDateParse(f), s, time.Local)
}

func ParseDateTimeFmt(f string) string {
	return _formatDateTimeParse(f)
}
