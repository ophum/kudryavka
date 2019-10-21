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

func req() {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"

		head := req.Header.Get("User-Agent")
		fmt.Println("====User-Agent====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("Accept")
		fmt.Println("====Accept====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("Accept-Language")
		fmt.Println("====Accept-Language====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("Accept-Encoding")
		fmt.Println("====Accept-Encofing====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("Referer")
		fmt.Println("====Referer====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("Connection")
		fmt.Println("====Connection====")
		fmt.Printf("%s\n\n",head)

		head = req.Header.Get("Upgrade-Insecure-Requests")
		fmt.Println("====Upgrade-Insecure-Requests====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("If-Modified-Since")
		fmt.Println("====If-Modified-Since====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("If-None-Match")
		fmt.Println("====If-None-Match====")
		fmt.Printf("%s\n",head)

		head = req.Header.Get("Cache-Control")
		fmt.Println("====Cache-Control====")
		fmt.Printf("%s\n",head)

		origin, _ := url.Parse("http://localhost:12345")
		req.URL.Host = origin.Host
	}

	modres := func(res *http.Response) error {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		body := buf.String()
		//fmt.Println("=====RESPONSE BODY=====")
		//fmt.Println(body)
		fmt.Printf("--------------------------------------------\n\n")
		res.Body = ioutil.NopCloser(strings.NewReader(body))
		return nil
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modres}

	fmt.Println("Server running...")
	http.ListenAndServe(":10000", proxy)
}


