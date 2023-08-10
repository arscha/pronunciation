package util

import (
	"net/http"

	"golang.org/x/net/proxy"
)

func SetProxy(addr string) (*http.Client, error) {
	socksProxy, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{Dial: socksProxy.Dial}

	client := &http.Client{
		Transport: transport,
	}

	return client, nil
}
