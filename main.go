package main

// *chello* is a small command line tool which can interact with services which
// expose it's api via [hello][]. a sample service to toy around with is
// [hello_pingpong][].
//
// [hello]: https://github.com/travelping/hello
// [hello_pingpong]: https://github.com/liveforeverx/hello_pingpong%

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	var cli = struct {
		id       string
		addr     string
		encoding string
		indent   bool
		echo     bool
	}{
		addr:     "127.0.0.1:26000",
		encoding: "jsonrpc2",
	}

	flag.StringVar(&cli.id, "id", cli.id, "id to send with requests")
	flag.StringVar(&cli.encoding, "enc", cli.encoding, "which encoding to use")
	flag.StringVar(&cli.addr, "addr", defaultAddr(cli.addr),
		fmt.Sprintf("addr to use for transport (%s)", availableTransports()))
	flag.BoolVar(&cli.indent, "indent", cli.indent, "indent jsonrpc answer")
	flag.BoolVar(&cli.echo, "echo", cli.echo, "show what gets send (on stderr)")
	flag.Parse()

	var request = NewJSONRPCRequest(cli.id)
	if err := request.setMethodAndParams(flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	payload, _ := request.toJSON()
	if cli.echo {
		fmt.Fprintln(os.Stderr, payload)
	}

	transport, err := newTransport(cli.addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating transport: %v\n", err)
		os.Exit(1)
	}

	if err := transport.Connect(cli.addr); err != nil {
		fmt.Fprintf(os.Stderr, "error connecting: %v\n", err)
		os.Exit(1)
	}
	defer transport.Close()

	reply, err := transport.Request(payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error request: %v\n", err)
		os.Exit(1)
	}

	if cli.indent {
		var buf = bytes.NewBuffer(nil)
		json.Indent(buf, []byte(reply), "", "  ")
		reply = buf.String()
	}

	fmt.Println(reply)
}
