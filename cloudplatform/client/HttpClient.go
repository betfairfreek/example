package client

import "net/http"

func HttpClient() *http.Client {
	return &http.Client{}
}
