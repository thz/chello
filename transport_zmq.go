package main

import (
	zmq "github.com/pebbe/zmq2"
)

func init() {
	registerTransport("zmq", func() (helloTransport, error) {
		return NewZmqSocket()
	})
}

type zmqSocket struct {
	*zmq.Socket
}

func NewZmqSocket() (*zmqSocket, error) {
	var sock, err = zmq.NewSocket(zmq.REQ)
	if err == nil {
		return &zmqSocket{sock}, nil
	}
	return nil, err
}

func (s *zmqSocket) Connect(addr string) error {
	return s.Socket.Connect(addr)
}

func (s *zmqSocket) Close() error {
	return s.Socket.Close()
}

func (s *zmqSocket) Request(req string) (string, error) {
	var (
		reply []string
		err   error
	)

	if _, err = s.Socket.SendMessage(req); err != nil {
		return "", err
	}
	if reply, err = s.Socket.RecvMessage(0); err != nil {
		return "", err
	}

	return reply[0], nil
}
