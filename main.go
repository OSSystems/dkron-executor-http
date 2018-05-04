package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
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
		fmt.Println(err)
		os.Exit(1)
	}

	args := &Arguments{}

	if err = json.Unmarshal(stdin, args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := doRequest(args)
	if err != nil {
		if len(out) > 0 {
			fmt.Println(string(out))
		} else {
			fmt.Println(err)
		}

		os.Exit(1)
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

	if resp.StatusCode >= 300 {
		return out, errors.Errorf("response status code: %d", resp.StatusCode)
	}

	return out, nil
}
