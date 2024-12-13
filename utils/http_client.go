package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var Client = &http.Client{}

func RequestGET(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Check for non-2xx status codes
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to get data, status code: %d", res.StatusCode)
	}

	return body, nil
}

func RequestPOST(url string, headers map[string]string, dataPayload io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, dataPayload)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to post data, status code: %d", res.StatusCode)
	}

	return body, nil
}
