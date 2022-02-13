package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func doRequest(url, method string, headers map[string]string, rawBody interface{}) ([]byte, error) {
	var body io.Reader = nil
	if rawBody != nil {
		switch t := rawBody.(type) {
		case string:
			body = strings.NewReader(t)
		case []byte:
			body = bytes.NewReader(t)
		default:
			data, err := json.Marshal(rawBody)
			if err != nil {
				err = fmt.Errorf("body 序列化失败: %v", err)
				return nil, err
			}
			body = bytes.NewReader(data)
		}
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()
	responseData := make([]byte, 0)
	responseData, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 0 {
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			err = errors.New(resp.Status)
		}
	}
	return responseData, err
}

func DoPost(url string, headers map[string]string, rawBody interface{}) ([]byte, error) {
	return doRequest(url, "POST", headers, rawBody)
}
