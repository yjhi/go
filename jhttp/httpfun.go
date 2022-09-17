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
package jhttp

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (string, error) {
	resp, err1 := http.Get(url)

	if err1 != nil {
		return "", err1
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return "", err2
	}
	return string(body), nil
}

func PostJson(url string, d string) (string, error) {
	return Post(url, "application/json", d)
}

func Post(url string, t string, d string) (string, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(d))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return "", err2
	}
	return string(body), nil

}
