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

package jhttp

import (
	"net/http"
)

///Add By yjh 2021-06-17
type Http struct {
	Client    *http.Client
	Error     string
	Request   *http.Request
	httpError error
	isOk      bool
}

func New(url string, timeoutSecond int, way string) *Http {
	return _buildHttp(url, timeoutSecond, way)
}

func Create(url string, timeoutSecond int, way string) *Http {
	return _buildHttp(url, timeoutSecond, way)
}
func CreateGet(url string, timeoutSecond int) *Http {

	return Create(url, timeoutSecond, "GET")
}

func CreatePOST(url string, timeoutSecond int) *Http {

	return Create(url, timeoutSecond, "POST")
}

func (h *Http) AddCookie(name string, value string, path string, domain string) *Http {

	if h.isOk {
		h.Request.AddCookie(&http.Cookie{
			Name:   name,
			Value:  value,
			Path:   path,
			Domain: domain,
		})
	}

	return h
}
func (h *Http) AddHeader(name string, value string) *Http {
	if h.isOk {
		h.Request.Header.Set(name, value)
	}

	return h
}

func (h *Http) SetHeader(name string, value string) *Http {
	return h.AddHeader(name, value)
}

func (h *Http) SendRequest() (string, error) {
	return h._do()
}
func (h *Http) SendRequestWithoutBody() (int, error) {
	return h._doWithoutBody()
}
func (h *Http) SendRequestWithResp() (*HttpResp, error) {
	return h._doWithResp()
}
