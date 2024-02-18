package http

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			MaxIdleConns:      10,
			IdleConnTimeout:   30 * time.Second,
		},
		Timeout: 10 * time.Second,
	}
)

// SendRequest 发送请求
func SendRequest(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
