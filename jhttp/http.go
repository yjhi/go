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
	"time"

	"gitee.com/yjhi/go/jerrors"
)

type Http struct {
	Client  *http.Client
	Error   string
	Request *http.Request
	IsOk    bool
}

func Create(url string, secondtimeout int, way string) *Http {
	return New(url, secondtimeout, way)
}

func New(url string, secondtimeout int, way string) *Http {

	h := &Http{
		Client:  nil,
		Request: nil,
		Error:   "",
		IsOk:    false,
	}

	h.Client = &http.Client{
		Timeout: time.Duration(secondtimeout) * time.Second,
	}

	var err error

	h.Request, err = http.NewRequest(way, url, nil)

	if err != nil {
		h.Error = err.Error()
		return h
	}

	h.IsOk = true

	return h
}

func CreateGet(url string, secondtimeout int) *Http {

	return Create(url, secondtimeout, "GET")
}

func CreatePOST(url string, secondtimeout int) *Http {

	return Create(url, secondtimeout, "POST")
}

func (h *Http) AddCookie(name string, value string, path string, domain string) *Http {

	if h.IsOk {
		h.Request.AddCookie(&http.Cookie{
			Name:   name,
			Value:  value,
			Path:   path,
			Domain: domain,
		})
	}

	return h
}

func (h *Http) SetHeader(name string, value string) *Http {
	if h.IsOk {
		h.Request.Header.Set(name, value)
	}

	return h
}

func (h *Http) Do() (string, error) {

	if h.IsOk {
		resp, err := h.Client.Do(h.Request)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		bodyContent, errret := ioutil.ReadAll(resp.Body)

		if errret != nil {
			return "", errret
		}

		return string(bodyContent), nil

	} else {

		return "", jerrors.NewError("Http 初始化失败")
	}

}

func (h *Http) DoWithoutBody() error {

	if h.IsOk {
		_, err := h.Client.Do(h.Request)
		if err != nil {
			return err
		}
		return nil

	} else {

		return jerrors.NewError("Http 初始化失败")
	}

}

func (h *Http) DoWithResp() (*HttpResp, error) {

	s := &HttpResp{
		Response: nil,
	}

	var err error

	s.Response, err = h.Client.Do(h.Request)

	if err != nil {
		return nil, err
	}

	return s, nil
}
