package requests

import (
	"net/http"
)

func request(method string, url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	return res, nil
}

func Get(url string, headers map[string]string) (*http.Response, error) {
	return request("GET", url, nil)
}

func Options(url string, headers map[string]string) (*http.Response, error) {
	return request("OPTIONS", url, nil)
}

func Head(url string, headers map[string]string) (*http.Response, error) {
	return request("HEAD", url, nil)
}

func Post(url string, headers map[string]string) (*http.Response, error) {
	return request("POST", url, headers)
}

func Put(url string, headers map[string]string) (*http.Response, error) {
	return request("PUT", url, headers)
}

func Patch(url string, headers map[string]string) (*http.Response, error) {
	return request("PATCH", url, headers)
}

func Delete(url string) (*http.Response, error) {
	return request("DELETE", url, nil)
}
