package gzlm

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"time"
)

var defaultClient *resty.Client

func Request() *resty.Request {
	if defaultClient != nil {
		return defaultClient.R()
	}

	defaultClient = resty.New().
		SetRetryResetReaders(true).
		SetTimeout(time.Second * 30).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return defaultClient.R()
}

func RequestWithTimeout(t time.Duration, retryCount int) *resty.Request {
	return resty.New().
		SetRetryCount(retryCount).
		SetRetryResetReaders(true).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTimeout(t).R()
}
