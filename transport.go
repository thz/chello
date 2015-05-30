package main

import (
	"fmt"
	"strings"
)

type helloTransport interface {
	Connect(addr string) error
	Close() error
	Request(req string) (reply string, err error)
}

type transGenerator func() (helloTransport, error)

var transporter map[string]transGenerator = make(map[string]transGenerator)

func registerTransport(id string, gen transGenerator) {
	transporter[id] = gen
}

func availableTransports() string {
	var transports = make([]string, 0, len(transporter))
	for id := range transporter {
		transports = append(transports, id)
	}
	return strings.Join(transports, ",")
}

func NewTransport(id string) (helloTransport, error) {
	if generate, ok := transporter[id]; ok {
		return generate()
	}
	return nil, fmt.Errorf("invalid transport %q", id)
}
