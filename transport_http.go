package main

import (
    "net/http"
    "bytes"
    "io"
    "strings"
)

func init() {
    registerTransport("http", func() (helloTransport, error) {
        return &Saver{}, nil
    })
}

type Saver struct {
    addr string
}

func (s *Saver) Connect(addr string) error {
    s.addr = addr
    return nil
}

func (s *Saver) Close() error {
    return nil
}

func (s *Saver) Request(req string) (string, error) {
    request := strings.NewReader(req)
    reply := bytes.NewBuffer(nil)
    resp, err := http.Post(s.addr, "application/json", request)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    _, err = io.Copy(reply, resp.Body)
    return reply.String(), err
}
