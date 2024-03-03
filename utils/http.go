package utils

import (
	"net"
	"net/http"
	"time"
)

var (
	dialer = &net.Dialer{
		Timeout:   time.Second,
		KeepAlive: 60 * time.Second,
	}

	transport = &http.Transport{
		DialContext:         dialer.DialContext,
		MaxIdleConnsPerHost: 50,
		MaxConnsPerHost:     50,
		IdleConnTimeout:     90 * time.Second,
	}

	Client = &http.Client{
		Timeout:   20 * time.Second,
		Transport: transport,
	}
)
