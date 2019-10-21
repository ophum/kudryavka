package main

import (
	"net/http/httputil"
	"net/http"
	"net/url"
	"fmt"
	"bytes"
	"io/ioutil"
	"strings"
)

func main() {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		body := req.Body
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
		res.Body = ioutil.NopCloser(strings.NewReader(body))
		return nil
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modres}

	fmt.Println("Server running...")
	http.ListenAndServe(":8000", proxy)
}


