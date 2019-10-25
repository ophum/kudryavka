package main

import "C"
import (
	"fmt"
	"unsafe"
	"regexp"
//	"os"

	"github.com/ophum/kudryavka/waf/gate"
)

var pattern = []string{	`.*<script.*`,
												`.*<style.*`,
												`.*javascript.*`,
												`.*vbscript.*`,
												`.*about.*`,
												`.*expression.*`,
												`.*<link.*`,
												`.*&{.*`,
												`.*<body.*onload=.*`,
												`.*onMouseOver.*`,
												`.*onClick.*`,
												`.*document.cookie.*`,
												`.*%0d.*`,
												`.*%0a.*`,
												`.*Microsoft.XMLHTTP.*`,
											}

//export check
func check(args unsafe.Pointer) {
	req := (*gate.CheckList)(args)
	//fmt.Println(req.Body)

	for i:=0; i<len(pattern); i++ {
		re := regexp.MustCompile(pattern[i])
		if re.Copy().MatchString(req.Body){
			fmt.Printf("match: \n  %s\n", pattern[i])
		}else{
			//fmt.Println("no match")
			continue
		}
	}
}

func main() {
}
