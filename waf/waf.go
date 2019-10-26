package waf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/ophum/kudryavka/waf/gate"
)

type Waf struct {
	server *url.URL
	input  *gate.Gates
	output *gate.Gates
}

func NewWaf(server *url.URL) *Waf {
	i := gate.NewGates()
	o := gate.NewGates()

	return &Waf{
		server: server,
		input:  i,
		output: o,
	}
}

func (w *Waf) Serve() {
	director := func(req *http.Request) {
		req.URL.Scheme = w.server.Scheme
		r := gate.CheckList{}

		err := w.input.CheckAll(r)
		if err != nil {
			// error
		}
		req.URL.Host = w.server.Host
	}

	modres := func(res *http.Response) error {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		body := buf.String()
		fmt.Println("=====RESPONSE BODY=====")
		fmt.Println(body)
		b, _ := json.Marshal(res)
		fmt.Println("json")
		fmt.Println(string(b))
		r := gate.CheckList{}
		err := w.output.CheckAll(r)
		if err != nil {
			// error
		}
		res.Body = ioutil.NopCloser(strings.NewReader(body))
		return nil
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modres}

	fmt.Println("Server running...")
	http.ListenAndServe(":8000", proxy)
}
