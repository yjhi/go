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
package jutils

import (
	"bytes"
	"math"
)

const encodeStd = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type Encoding struct {
	encodeMap [62]byte
	decodeMap map[byte]int
}

func NewEncoding() *Encoding {

	e := &Encoding{}

	copy(e.encodeMap[:], encodeStd)

	e.decodeMap = make(map[byte]int)

	for i, v := range e.encodeMap {
		e.decodeMap[v] = i
	}

	return e

}

func (en *Encoding) ToBase62(indata int64) string {
	var outStr bytes.Buffer

	last := indata

	for last != 0 {

		i := last % 62

		outStr.WriteByte(en.encodeMap[i])
		last = (last - i) / 62

	}

	s := outStr.String()

	return s
}

func (en *Encoding) FromBase62(indata string) int64 {
	var ret int64
	slen := len(indata)

	for i := 0; i < slen; i++ {
		ret += int64(en.decodeMap[indata[i]] * int(math.Pow(62, float64(i))))
	}

	return ret

}
