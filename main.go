package main

import (
	"github.com/ophum/kudryavka/waf/gate"
	"github.com/ophum/kudryavka/waf/gate/pattern"
)

func main() {
	//u, _ := url.Parse("http://localhost:8081")
	//kudryavka := waf.NewWaf(u)
	//kudryavka.Serve()
	gates := gate.NewGates()
	pgate := pattern.NewPatternGate()
	pgate.Append(pattern.Method, `GET`)
	gates.Append("pattern get", pgate)
	req := gate.CheckList{
		Method: "GET",
	}
	gates.CheckAll(req)
	//	dgate, err := dynamic.NewDynamicGate("./hello.so")
	//	if err != nil {
	//		return
	//	}
	//	dgate2, err := dynamic.NewDynamicGate("./gostruct.so")
	//	if err != nil {
	//		return
	//	}
	//
	//	gates.Append("helli", dgate)
	//	gates.Append("gostruct", dgate2)
	//	req := gate.CheckList{
	//		Method: "GET",
	//	}
	//
	//	gates.CheckAll(req)
}
