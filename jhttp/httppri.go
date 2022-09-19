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
	"io/ioutil"
	"net/http"
	"time"
)

func _buildHttp(url string, timeoutSecond int, way string) *Http {

	h := &Http{
		Client:  nil,
		Request: nil,
		Error:   "",
		isOk:    false,
	}

	h.Client = &http.Client{
		Timeout: time.Duration(timeoutSecond) * time.Second,
	}

	h.Request, h.httpError = http.NewRequest(way, url, nil)

	if h.httpError != nil {
		h.Error = h.httpError.Error()
		return h
	}

	h.isOk = true

	return h
}

func (h *Http) _do() (string, error) {

	if h.isOk {
		resp, err := h.Client.Do(h.Request)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		bodyContent, errRet := ioutil.ReadAll(resp.Body)

		if errRet != nil {
			return "", errRet
		}

		return string(bodyContent), nil

	} else {

		return "", h.httpError
	}

}

func (h *Http) _doWithoutBody() (int, error) {

	if h.isOk {
		resp, err := h.Client.Do(h.Request)
		if err != nil {
			return -1, err
		}

		return resp.StatusCode, nil

	} else {

		return -1, h.httpError
	}

}

func (h *Http) _doWithResp() (*HttpResp, error) {

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
