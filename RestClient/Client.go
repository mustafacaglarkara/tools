package RestClient

import (
	"bytes"
	"net/http"
)

type Client struct {
}

// HTTP istekleri için genel bir fonksiyon oluşturuyoruz.
func sendRequest(method string, url string, headers http.Header, auth string, data []byte) (*http.Response, error) {

	// İstek yapmak için bir httpClient oluşturuyoruz.
	httpClient := &http.Client{}

	// HTTP isteğini hazırlıyoruz.
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// Header bilgilerini ekliyoruz.
	if headers != nil {
		req.Header = headers
	}

	// Auth bilgisini ekliyoruz.
	if auth != "" {
		req.Header.Add("Authorization", auth)
	}

	// İsteği gönderiyoruz ve cevabı dönüyoruz.
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// POST isteği için bir fonksiyon oluşturuyoruz.
func (c *Client) Post(url string, headers http.Header, auth string, data []byte) (*http.Response, error) {
	return sendRequest("POST", url, headers, auth, data)
}

// PUT isteği için bir fonksiyon oluşturuyoruz.
func (c *Client) Put(url string, headers http.Header, auth string, data []byte) (*http.Response, error) {
	return sendRequest("PUT", url, headers, auth, data)
}

// GET isteği için bir fonksiyon oluşturuyoruz.
func (c *Client) Get(url string, headers http.Header, auth string) (*http.Response, error) {
	return sendRequest("GET", url, headers, auth, nil)
}
