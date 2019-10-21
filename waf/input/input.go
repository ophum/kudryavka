package input

import "net/http"
import "github.com/ophum/waf/gate"

type Gate interface {
	Check()
}

type XSSGate struct {
	headers []string
}

func (x *XSSGate) Check() {
	x.headers
	print("helllo")
}

type CSRFGate struct {
	sss string
}

func (c *CSRFGate) Check() {
	c.sss 
	print("worla")
}
type Input struct {
	// チェック項目を持つ
	Gates []gate.Gate
}

func NewInput() *Input {
	return &Input{}
}

// チェック項目を追加する
func (i *Input) Append() {

}

func (i *Input) Check(req *http.Request) {
	// チェック項目を見ていってチェック
	for _, v := range i.Gates {
		v.Check()
	}
}

func main() {

	in := Input{
		Gates: []Gate{
			&XSSGate{},
			&CSRFGate{},

		}
	}

	in.Append(&AAAGate{})
	in.Gates[0].Check() // XSSGate.Check()
	in.Gates[1].Check() // CSRFGate.Check()
}
