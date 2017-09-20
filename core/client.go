package core

import (
	"time"
	"net/http"
)

func Client() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       1000,
		MaxIdleConnsPerHost: 1000,
		IdleConnTimeout:    30 * time.Second,
	}
	return &http.Client{Transport: tr}
}