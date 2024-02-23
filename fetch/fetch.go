package fetch

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"time"
)

func addHTTPHeaders(req *http.Request) {
	for k, v := range HEADERS {
		req.Header.Add(k, v)
	}
}

func commonClient() *http.Client {
	trans := &http.Transport{
		MaxConnsPerHost: 10,
		MaxIdleConns:    90,
	}
	client := http.Client{
		Timeout:   10 * time.Second,
		Transport: trans,
	}
	return &client
}

func Do(target string, ungzip bool) (string, error) {
	client := commonClient()

	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return "", err
	}
	addHTTPHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http status code: %v", resp.StatusCode)
	}

	result := resp.Body
	if ungzip {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			if gr != nil {
				gr.Close()
			}
			return "", err
		}
		defer gr.Close()
		result = gr
	}

	raw, err := io.ReadAll(result)
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer(raw)
	return buff.String(), nil
}
