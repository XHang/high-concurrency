package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"spike/log"
	"testing"
)

func GetTest(t *testing.T, urlStr string, param map[string]string) []byte {
	urlParam := url.Values{}
	for k, v := range param {
		urlParam.Add(k, v)
	}
	if len(param) > 0 {
		urlStr += "?" + urlParam.Encode()
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		t.Error(err)
		return nil
	}
	rr := httptest.NewRecorder()
	Engine.ServeHTTP(rr, req)
	bytes, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
		return nil
	}
	if rr.Result().StatusCode != 200 {
		t.Errorf("请求发生错误，状态码【%v】不是为200", rr.Result().StatusCode)
		return bytes
	}

	return bytes
}
func init() {
	Init()
	log.Init()
}
