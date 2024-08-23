package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 执行HTTP POST请求的函数
func postRequest(ctx context.Context, url string, req interface{}, resp interface{}) (err error) {
	var (
		bodydata []byte
		reqhttp  *http.Request
		resphttp *http.Response
	)
	if bodydata, err = json.Marshal(req); err != nil {
		fmt.Printf("failed to marshal request data: %s", err.Error())
		return
	}

	//创建HTTP请求
	reqhttp, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodydata))
	if err != nil {
		fmt.Printf("failed to create request: %s", err.Error())
		return
	}
	reqhttp.Header.Set("Content-Type", "application/json")
	// 执行HTTP请求
	client := &http.Client{}
	resphttp, err = client.Do(reqhttp)
	if err != nil {
		fmt.Printf("failed to do request: %s", err.Error())
		return
	}
	defer resphttp.Body.Close()

	// 读取响应数据
	body, err := io.ReadAll(resphttp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %s", err.Error())
		return
	}

	// 检查HTTP状态码
	if resphttp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code: %d, body: %s", resphttp.StatusCode, string(body))
		return
	}

	// 解析响应数据
	err = json.Unmarshal(body, resp)
	if err != nil {
		fmt.Printf("failed to unmarshal response err: %v json:%s", err, string(body))
		return
	}
	return
}
