package gate

import (
	"fmt"
	"log"
)

type Gate interface {
	Check(CheckList) error
}

// ポインター渡せないので、基本構造体はばらしていく方向で
type CheckList struct {
	Method string
	//URL        *url.URL
	Proto      string
	ProtoMajor int
	ProtoMinor int

	//Header http.Header
	//Body             io.ReadCloser
	ContentLength int64
	//TransferEncoding []string
	Host string
	//From     url.Values
	//PostForm url.Values
	//MultipartForm    *multipart.Form
	//Trailer    http.Header
	RemoteAddr string
	RequestURI string
}

type Gates struct {
	gates map[string]Gate
}

func NewGates() *Gates {
	return &Gates{
		gates: map[string]Gate{},
	}
}

func (g *Gates) Append(name string, gate Gate) {
	g.gates[name] = gate
}

// goroutineで回したい
// エラー処理
func (g *Gates) CheckAll(list CheckList) error {
	for name, f := range g.gates {
		log.Printf("Checing %s...\n", name)
		err := f.Check(list)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
