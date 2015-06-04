package main

import (
	"fmt"
	"net/url"
	"sort"
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

func transportList() []string {
	var transports = make([]string, 0, len(transporter))
	for id := range transporter {
		transports = append(transports, id)
	}
	sort.StringSlice(transports).Sort()
	return transports
}

func availableTransports() string {
	var transports = transportList()
	return strings.Join(transports, ",")
}

func firstTransport() string {
	var transports = transportList()
	return transports[0]
}

func defaultAddr(addr string) string {
	return fmt.Sprintf("%s://%s", firstTransport(), addr)
}

func newTransport(addr string) (helloTransport, error) {
	var uri, err = url.Parse(addr)
	if err != nil {
		return nil, err
	}
	if generate, ok := transporter[uri.Scheme]; ok {
		return generate()
	}
	return nil, fmt.Errorf("invalid transport %q", uri.Scheme)
}
