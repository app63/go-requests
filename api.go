package requests

import (
	"net/http"
	"strings"
)

func request(method string, url string, data strings.NewReader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, *data)

	res, _ := client.Do(req)
	defer req.Body.Close()

	return res, nil
}

func Get(url string) (*http.Response, error) {
	return request("GET", url, nil)
}

func Options(url string) (*http.Response, error) {
	return request("OPTIONS", url, nil)
}

func Head(url string) (*http.Response, error) {
	return request("HEAD", url, nil)
}

func Post(url string, data strings.NewReader) (*http.Response, error) {
	return request("POST", url, data)
}

func Put(url string, data strings.NewReader) (*http.Response, error) {
	return request("PUT", url, data)
}

func Patch(url string, data strings.NewReader) (*http.Response, error) {
	return request("PATCH", url, data)
}

func Delete(url string) (*http.Response, error) {
	return request("DELETE", url, nil)
}
