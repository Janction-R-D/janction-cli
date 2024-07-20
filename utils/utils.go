package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

func GetWithTimeout(reqUrl string, header map[string]string, params url.Values, timeout time.Duration) ([]byte, error) {
	var res []byte
	rUrl, err := url.Parse(reqUrl)
	if err != nil {
		return nil, errors.New("url parse error")
	}
	rUrl.RawQuery = params.Encode()
	urlPath := rUrl.String()

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return res, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	HttpClient := &http.Client{
		Timeout: timeout,
	}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	res, err = io.ReadAll(resp.Body)
	return res, err
}

func PostWithTimeout(url string, body []byte, header map[string]string, timeout time.Duration) ([]byte, error) {
	var res []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return res, err
	}
	HttpClient := &http.Client{
		Timeout: timeout,
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	res, err = io.ReadAll(resp.Body)
	return res, err
}
