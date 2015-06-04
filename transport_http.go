package main

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func init() {
	var gen = func() (helloTransport, error) {
		return &httpClient{}, nil
	}
	registerTransport("https", gen)
	registerTransport("http", gen)
}

type httpClient struct {
	addr string
}

func (c *httpClient) Connect(addr string) error {
	c.addr = addr
	return nil
}

func (c *httpClient) Close() error {
	return nil
}

func (c *httpClient) Request(req string) (string, error) {
	var request = strings.NewReader(req)
	var resp, err = http.Post(c.addr, "application/json", request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var reply = bytes.NewBuffer(nil)
	_, err = io.Copy(reply, resp.Body)
	return reply.String(), err
}
