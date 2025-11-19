package helper

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// RequestGet 发起 Get 请求
func RequestGet(url string, getParams string) ([]byte, error) {
	// 发送请求
	req, _ := http.NewRequest("POST", url+getParams, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, err
}

// RequestPost 发起 Post 请求
func RequestPost(url string, getParams string, requestBody interface{}) ([]byte, error) {
	// 封装 Post 参数
	jsonStr, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(jsonStr))

	// 发送请求
	req, _ := http.NewRequest("POST", url+getParams, payload)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, err
}

// RequestDelete 发起 Delete 请求
func RequestDelete(url string, getParams string) ([]byte, error) {
	// 发送请求
	req, _ := http.NewRequest("DELETE", url+getParams, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, err
}
