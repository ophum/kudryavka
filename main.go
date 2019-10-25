package main

import (
	"fmt"

	"github.com/ophum/kudryavka/waf/gate"
	"github.com/ophum/kudryavka/waf/gate/dynamic"
)
func main() {
	var pattern = []string{		`>"><script src=...>...</script>"`,
													 	`"><style type="text/javascript">...</style>`,
														`javascript: ...`,
														`expression( ...`,
														`<link rel="stylesheet" href="http://xxxx/bad.css=">`,
														`"><img src="" onError="...;">`,
														`a href="&{...};"`,
														`><body onload= ...`,
														`document.cookie`,
														`%0d%0a%0d%0a<script src="...">...`,
														`ActiveXObject("Microsoft.XMLHTTP")`,
													}

	//kudryavka := waf.NewWaf()
	//kudryavka.Serve()
	gates := gate.NewGates()
	dgate2, err := dynamic.NewDynamicGate("./xss.so")
	if err != nil {
		return
	}
	gates.Append("xss", dgate2)

	for i:=0; i<len(pattern); i++ {
		fmt.Printf("\n=========\nbody pattern : %s\n",pattern[i])
		req := gate.CheckList{
			Method: "GET",
			Body: pattern[i],
		}
		gates.CheckAll(req)
	}
}
