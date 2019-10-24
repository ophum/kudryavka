# modles - shared object

## example
```
package main

import "C"
import (
	"fmt"
	"unsafe"

	"github.com/ophum/kudryavka/waf/gate"
)

//export check
func check(args unsafe.Pointer) {
	req := (*gate.CheckList)(args)
	fmt.Println(req)
}

func main() {

}

```

checkという名前の関数で定義
```
//export check
func check(args unsafe.Pointer) {
```

引数をgate.CheckListにキャストし利用する
```
	req := (*gate.CheckList)(args)
}
```

## build
`test.go`をビルドし`test.so`を作成する
```
# go build -buildmode=c-shared -o test.so test.go
```
