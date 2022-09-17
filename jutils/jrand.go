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
	"math/rand"
	"time"
)

var IsInitRandSeed bool

func InitRandSeed() {
	if IsInitRandSeed != true {
		rand.Seed(time.Now().UnixNano())
		IsInitRandSeed = true

	}
}

func NewRandN(max int) func() int {
	InitRandSeed()
	return func() int {
		return rand.Intn(max + 1)
	}
}

func NewRandMN(min int, max int) func() int {
	InitRandSeed()
	return func() int {
		return rand.Intn(max+1) + min
	}
}

func NewRandN32(max int32) func() int32 {
	InitRandSeed()
	return func() int32 {
		return rand.Int31n(max + 1)
	}
}

func NewRandMN32(min int32, max int32) func() int32 {
	InitRandSeed()
	return func() int32 {
		return rand.Int31n(max+1) + min
	}
}

func NewRandN64(max int64) func() int64 {
	InitRandSeed()
	return func() int64 {
		return rand.Int63n(max + 1)
	}
}

func NewRandMN64(min int64, max int64) func() int64 {
	InitRandSeed()
	return func() int64 {
		return rand.Int63n(max+1) + min
	}
}

type RandStrUtil struct {
	strByte []byte
	count   int
}

// make rand str in strmap
// add by yjh
// 211110
func (r *RandStrUtil) Rand(count int) string {

	sbyte := make([]byte, 0)

	fRand := NewRandN(r.count - 1)

	for i := 0; i < count; i++ {
		sbyte = append(sbyte, r.strByte[fRand()])
	}

	return string(sbyte)
}

func BuildRandStr(str string) *RandStrUtil {
	r := &RandStrUtil{
		strByte: []byte(str),
		count:   len(str),
	}

	return r
}
