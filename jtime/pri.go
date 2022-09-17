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

import (
	"strings"
	"time"
)

func _formatDateParse(sf string) string {
	sf = strings.Replace(sf, "yyyy", "2006", 1)
	sf = strings.Replace(sf, "MM", "01", 1)
	sf = strings.Replace(sf, "dd", "02", 1)

	return sf
}

func _formatTimeParse(sf string) string {

	sf = strings.Replace(sf, "hh", "15", 1)
	sf = strings.Replace(sf, "mm", "04", 1)
	sf = strings.Replace(sf, "ss", "05", 1)

	return sf
}

func _formatDateTimeParse(sf string) string {
	sf = _formatDateParse(sf)
	sf = _formatTimeParse(sf)
	return sf
}

func _timeStr(f string) string {
	return time.Now().Format(f)
}
