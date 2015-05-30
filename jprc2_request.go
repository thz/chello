package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type jsonRPCRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	Id      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

func NewJSONRPCRequest(id string) *jsonRPCRequest {
	var req = new(jsonRPCRequest)
	req.JsonRPC = "2.0"
	req.Id = id
	if id == "" {
		req.Id = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return req
}

func (r *jsonRPCRequest) setMethodAndParams(args []string) error {
	switch {
	case len(args) == 0:
		return fmt.Errorf("missing 'method' (first argument)")
	case len(args) == 1:
		r.Method = args[0]
	case len(args) > 1:
		r.Method = args[0]

		if strings.ContainsRune(args[1], '=') {
			// if first argument conatins a '=' i assume
			// a list of key-value pairs, forming a map
			// or an 'object' in javascript terms
			r.Params = kvPairsToMap(args[1:])

		} else if len(args[1]) > 1 && args[1][0] == '{' {
			var obj, err = stringToJSObject(args[1])
			if err != nil {
				return err
			}
			r.Params = obj
		} else {
			r.Params = args[1:]
		}
	}
	return nil
}

func (r *jsonRPCRequest) toJSON() (string, error) {
	var buf = bytes.NewBuffer(nil)
	var encoder = json.NewEncoder(buf)
	if err := encoder.Encode(r); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func kvPairsToMap(pairs []string) map[string]interface{} {
	var m = make(map[string]interface{})
	for _, elem := range pairs {
		var (
			parts = strings.SplitN(elem, "=", 2)
			key   = parts[0]
			val   interface{}
		)

		if len(parts) == 2 {
			// if the value can be parsed as a float, we assume
			// a json.Number
			if _, err := strconv.ParseFloat(parts[1], 64); err == nil {
				val = json.Number(parts[1])
			} else {
				val = parts[1]
			}
		}
		m[key] = val
	}
	return m
}

func stringToJSObject(s string) (interface{}, error) {
	var obj interface{}
	var decoder = json.NewDecoder(strings.NewReader(s))
	if err := decoder.Decode(&obj); err != nil {
		return nil, err
	}
	return obj, nil
}
