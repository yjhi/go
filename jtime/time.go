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

package jtime

import "time"

// parse time utils
// add by yjh 211124
type TimeUtils struct {
	FormatDate     string
	FormatTime     string
	FormatDateTime string
}

// parse time
// add by yjh 211124
func (t TimeUtils) Time() string {
	return _timeStr(t.FormatTime)
}

// parse date
//a dd by yjh 211124
func (t TimeUtils) Date() string {
	return _timeStr(t.FormatDate)
}

// parse datetime
// add by yjh 211124
func (t TimeUtils) DateTime() string {
	return _timeStr(t.FormatDateTime)
}

// parse datetime with fmt
// add by yjh 211124
func (t TimeUtils) TimeWithFmt(f string) string {
	return _timeStr(_formatTimeParse(f))
}
func (t TimeUtils) DateWithFmt(f string) string {
	return _timeStr(_formatDateParse(f))
}
func (t TimeUtils) DateTimeWithFmt(f string) string {
	return _timeStr(_formatDateTimeParse(f))
}

// count time
// add by yjh 211125
type TimeCount struct {
	_startTime time.Time
	_stopTime  time.Time
}

func (t *TimeCount) Start() {
	t._startTime = time.Now()
}

func (t *TimeCount) Minutes() float64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Minutes()
}
func (t *TimeCount) Seconds() float64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Seconds()
}
func (t *TimeCount) Hours() float64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Hours()
}

func (t *TimeCount) Microseconds() int64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Microseconds()
}
func (t *TimeCount) Milliseconds() int64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Milliseconds()
}
