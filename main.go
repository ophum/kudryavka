package main

import (
	"github.com/ophum/kudryavka/waf/gate"

	"github.com/ophum/kudryavka/waf/gate/dynamic"
)

func main() {
	//kudryavka := waf.NewWaf()
	//kudryavka.Serve()
	gates := gate.NewGates()
	dgate, err := dynamic.NewDynamicGate("./hello.so")
	if err != nil {
		return
	}
	dgate2, err := dynamic.NewDynamicGate("./gostruct.so")
	if err != nil {
		return
	}

	gates.Append("helli", dgate)
	gates.Append("gostruct", dgate2)
	req := gate.CheckList{
		Method: "GET",
	}

	gates.CheckAll(req)
}
