package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq2"
)

type zmqSocket struct {
	*zmq.Socket
}

func init() {
	var generator = func() helloTransport {
		var sock, err = zmq.NewSocket(zmq.DEALER)

		if err != nil {
			panic(fmt.Errorf("error: %v\n", err))
		}
		return &zmqSocket{Socket: sock}
	}
	registerTransport("zmq", generator)
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

	// NOTE: for whatever reasons, we need to send an empty
	// frame first.
	if _, err = s.Socket.SendMessage("", req); err != nil {
		return "", err
	}

	if reply, err = s.Socket.RecvMessage(0); err != nil {
		return "", err
	}

	// NOTE: we also get 2 frames, the first one is empty
	if len(reply) != 2 {
		return "", fmt.Errorf("expected 2 frames, got %d", len(reply))
	}

	return reply[1], nil
}
