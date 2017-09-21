package core

import (
	"time"
	"net/http"
)

var Client = client()

func client() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       1000,
		MaxIdleConnsPerHost: 1000,
		IdleConnTimeout:    30 * time.Second,
	}
	return &http.Client{Transport: tr}
}