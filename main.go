package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/ophum/waf/input"
	"github.com/ophum/waf/output"
)

func AAA(req *http.Request) error {

	return nil
}
func main() {
	in := input.NewInput()
	out := output.NewOutput()
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		body := req.Body
		err := AAA(req)
		if err != nil {

		}
		in.Check(req)
		fmt.Println("=====BODY=====")
		fmt.Println(body)

		origin, _ := url.Parse("http://localhost:8081")
		req.URL.Host = origin.Host
	}

	modres := func(res *http.Response) error {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		body := buf.String()
		fmt.Println("=====RESPONSE BODY=====")
		fmt.Println(body)
		out.Check(res)
		res.Body = ioutil.NopCloser(strings.NewReader(body))
		return nil
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modres}

	fmt.Println("Server running...")
	http.ListenAndServe(":8000", proxy)
}
