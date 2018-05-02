package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Arguments struct {
	URL    string            `json:"url"`
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   []byte            `json:"body"`
}

func main() {
	stdin, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}

	args := &Arguments{}

	if err = json.Unmarshal(stdin, args); err != nil {
		panic(err)
	}

	out, err := doRequest(args)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(out))
}

func doRequest(args *Arguments) ([]byte, error) {
	cli := &http.Client{}

	req, err := http.NewRequest(args.Method, args.URL, bytes.NewReader(args.Body))
	if err != nil {
		return nil, err
	}

	for k, v := range args.Header {
		req.Header.Add(k, v)
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return out, nil
}
