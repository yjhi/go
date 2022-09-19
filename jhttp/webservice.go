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
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gitee.com/yjhi/go/jerrors"
)

func _webService1(url string, fullbody string, errname string) (string, error) {

	res, err := http.Post(url, "text/xml; charset=UTF-8", strings.NewReader(fullbody))
	if nil != err {
		s := fmt.Sprintf("["+errname+"]Http Post Err:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	defer res.Body.Close()

	if 200 != res.StatusCode {

		s := fmt.Sprintf("["+errname+"]WebService Request Fail, Status Is: %d", res.StatusCode)

		return "", jerrors.NewError(s)
	}

	data, err := ioutil.ReadAll(res.Body)
	if nil != err {

		s := fmt.Sprintf("["+errname+"]Read Body err:%s", err)

		return "", jerrors.NewError(s)
	}
	return string(data), nil
}

func _webService(url string, fullbody string, errname string) (string, error) {

	trans := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	httpClient := &http.Client{
		Timeout:   time.Duration(60) * time.Second,
		Transport: trans,
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(fullbody))

	if err != nil {
		return "", jerrors.NewError(err.Error())
	}

	request.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(request)

	if nil != err {
		s := fmt.Sprintf("["+errname+"]Http Post Err:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	defer resp.Body.Close()

	if 200 != resp.StatusCode {

		s := fmt.Sprintf("["+errname+"]WebService Request Fail, Status Is: %d", resp.StatusCode)

		return "", jerrors.NewError(s)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {

		s := fmt.Sprintf("["+errname+"]Read Body err:%s", err)

		return "", jerrors.NewError(s)
	}
	return string(data), nil
}

///send a request of webservice
///version 1.1
func WebService11(url string, body string) (string, error) {
	reqBody := `<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
  xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
  xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"> 
  <soap:Body>` + body +
		`</soap:Envelope>`

	return _webService(url, reqBody, "Soap1.1")
}

///send a request of webservice
///version 1.2
func WebService12(url string, body string) (string, error) {
	reqBody := `<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
  xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
  xmlns:soap12="http://www.w3.org/2003/05/soap-envelope"> 
  <soap12:Body>` + body + `
  </soap12:Body> 
  </soap12:Envelope>`

	return _webService(url, reqBody, "Soap1.2")
}
