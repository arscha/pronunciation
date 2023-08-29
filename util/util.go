package util

import (
	"net/http"

	"github.com/joho/godotenv"
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

func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
