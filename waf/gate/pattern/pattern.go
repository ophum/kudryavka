package pattern

import (
	"fmt"
	"regexp"

	"github.com/ophum/kudryavka/waf/gate"
)

type PatternGate struct {
	// pattern file path
	path    string
	pattern []*Pattern
}

func NewPatternGate() *PatternGate {
	return &PatternGate{}
}

func (p *PatternGate) Check(list gate.CheckList) error {

	for _, v := range p.pattern {
		var err error
		switch v.target {
		case Method:
			err = v.Match(list.Method)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PatternGate) Append(target Target, raw string) {
	pat := NewPattern(target, raw)
	p.pattern = append(p.pattern, pat)
}

type Pattern struct {
	target Target
	reg    *regexp.Regexp
}

func NewPattern(target Target, raw string) *Pattern {
	pat := &Pattern{
		target: target,
		reg:    regexp.MustCompile(raw),
	}
	return pat
}

func (p *Pattern) Match(str string) error {
	if p.reg.Copy().MatchString(str) {
		return fmt.Errorf("Match illegal string...")
	}
	return nil
}
