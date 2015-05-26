package main

import "fmt"

type helloTransport interface {
	Connect(addr string) error
	Close() error
	Request(req string) (reply string, err error)
}

type genTransport func() helloTransport

var transporter map[string]genTransport = make(map[string]genTransport)

func registerTransport(id string, generator genTransport) {
	transporter[id] = generator
}

func NewTransport(id string) (helloTransport, error) {
	var generator, ok = transporter[id]
	if ok == true {
		return generator(), nil
	}
	return nil, fmt.Errorf("invalid transport %q", id)
}
