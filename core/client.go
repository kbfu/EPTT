package core

import (
	"net/http"
	"time"
	"crypto/tls"
)

var Client = client()

func client() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        10000,
		MaxIdleConnsPerHost: 10000,
		IdleConnTimeout:     30 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: tr}
}
